package binding

import (
	jsonpkg "encoding/json"
	"errors"
	"fmt"
	"github.com/21888/go-tagexpr-new"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/21888/go-tagexpr-new/ameda-loc"
	"github.com/tidwall/gjson"
)

const (
	specialChar = "\x07"
)

type paramInfo struct {
	fieldSelector  string
	structField    reflect.StructField
	tagInfos       []*tagInfo
	omitIns        map[in]bool
	bindErrFactory func(failField, msg string) error
	looseZeroMode  bool
	defaultVal     []byte
}

func (p *paramInfo) name(_ in) string {
	var name string
	for _, info := range p.tagInfos {
		if info.paramIn == json {
			name = info.paramName
			break
		}
	}
	if name == "" {
		return p.structField.Name
	}
	return name
}

func (p *paramInfo) getField(expr *tagexpr.TagExpr, initZero bool) (reflect.Value, error) {
	fh, found := expr.Field(p.fieldSelector)
	if found {
		v := fh.Value(initZero)
		if v.IsValid() {
			return v, nil
		}
	}
	return reflect.Value{}, nil
}

func (p *paramInfo) bindRawBody(info *tagInfo, expr *tagexpr.TagExpr, bodyBytes []byte) error {
	if len(bodyBytes) == 0 {
		if info.required {
			return info.requiredError
		}
		return nil
	}
	v, err := p.getField(expr, true)
	if err != nil || !v.IsValid() {
		return err
	}
	v = ameda_loc.DereferenceValue(v)
	switch v.Kind() {
	case reflect.Slice:
		if v.Type().Elem().Kind() != reflect.Uint8 {
			return info.typeError
		}
		v.Set(reflect.ValueOf(bodyBytes))
		return nil
	case reflect.String:
		v.Set(reflect.ValueOf(ameda_loc.UnsafeBytesToString(bodyBytes)))
		return nil
	default:
		return info.typeError
	}
}

func (p *paramInfo) bindPath(info *tagInfo, expr *tagexpr.TagExpr, pathParams PathParams) (bool, error) {
	if pathParams == nil {
		return false, nil
	}
	r, found := pathParams.Get(info.paramName)
	if !found {
		if info.required {
			return false, info.requiredError
		}
		return false, nil
	}
	return true, p.bindStringSlice(info, expr, []string{r})
}

func (p *paramInfo) bindQuery(info *tagInfo, expr *tagexpr.TagExpr, queryValues url.Values) (bool, error) {
	return p.bindMapStrings(info, expr, queryValues)
}

func (p *paramInfo) bindHeader(info *tagInfo, expr *tagexpr.TagExpr, header http.Header) (bool, error) {
	return p.bindMapStrings(info, expr, header)
}

func (p *paramInfo) bindCookie(info *tagInfo, expr *tagexpr.TagExpr, cookies []*http.Cookie) (bool, error) {
	var r []string
	for _, c := range cookies {
		if c.Name == info.paramName {
			r = append(r, c.Value)
		}
	}
	if len(r) == 0 {
		if info.required {
			return false, info.requiredError
		}
		return false, nil
	}
	return true, p.bindStringSlice(info, expr, r)
}

func (p *paramInfo) bindOrRequireBody(
	info *tagInfo, expr *tagexpr.TagExpr, bodyCodec codec, bodyString string,
	postForm map[string][]string, fileHeaders map[string][]*multipart.FileHeader, hasDefaultVal bool) (bool, error) {
	switch bodyCodec {
	case bodyForm:
		found, err := p.bindMapStrings(info, expr, postForm)
		if !found {
			return p.bindFileHeaders(info, expr, fileHeaders)
		}
		return found, err
	case bodyJSON:
		return p.checkRequireJSON(info, expr, bodyString, hasDefaultVal)
	case bodyProtobuf:
		// It has been checked when binding, no need to check now
		return true, nil
		// err := p.checkRequireProtobuf(info, expr, false)
		// return err == nil, err
	default:
		return false, info.contentTypeError
	}
}

func (p *paramInfo) checkRequireProtobuf(info *tagInfo, expr *tagexpr.TagExpr, checkOpt bool) error {
	if checkOpt && !info.required {
		v, err := p.getField(expr, false)
		if err != nil || !v.IsValid() {
			return info.requiredError
		}
	}
	return nil
}

func (p *paramInfo) checkRequireJSON(info *tagInfo, expr *tagexpr.TagExpr, bodyString string, hasDefaultVal bool) (bool, error) {
	var requiredError error
	if info.required { // only return error if it's a required field
		requiredError = info.requiredError
	} else if !hasDefaultVal {
		return true, nil
	}
	if !gjson.Get(bodyString, info.namePath).Exists() {
		idx := strings.LastIndex(info.namePath, ".")
		// There should be a superior but it is empty, no error is reported
		if idx > 0 && !gjson.Get(bodyString, info.namePath[:idx]).Exists() {
			return true, nil
		}
		return false, requiredError
	}
	v, err := p.getField(expr, false)
	if err != nil || !v.IsValid() {
		return false, requiredError
	}
	return true, nil
}

var fileHeaderType = reflect.TypeOf(multipart.FileHeader{})

func (p *paramInfo) bindFileHeaders(info *tagInfo, expr *tagexpr.TagExpr, fileHeaders map[string][]*multipart.FileHeader) (bool, error) {
	r, ok := fileHeaders[info.paramName]
	if !ok || len(r) == 0 {
		if info.required {
			return false, info.requiredError
		}
		return false, nil
	}
	v, err := p.getField(expr, true)
	if err != nil || !v.IsValid() {
		return true, err
	}
	v = ameda_loc.DereferenceValue(v)
	var elemType reflect.Type
	isSlice := v.Kind() == reflect.Slice
	if isSlice {
		elemType = v.Type().Elem()
	} else {
		elemType = v.Type()
	}
	var ptrDepth int
	for elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
		ptrDepth++
	}
	if elemType != fileHeaderType {
		return true, errors.New("parameter type is not (*)multipart.FileHeader struct or slice")
	}
	if len(r) == 0 || r[0] == nil {
		return true, nil
	}
	if !isSlice {
		v.Set(reflect.ValueOf(*r[0]))
		return true, nil
	}
	for _, fileHeader := range r {
		v.Set(reflect.Append(v, ameda_loc.ReferenceValue(reflect.ValueOf(fileHeader), ptrDepth-1)))
	}
	return true, nil
}

func (p *paramInfo) bindMapStrings(info *tagInfo, expr *tagexpr.TagExpr, values map[string][]string) (bool, error) {
	r, ok := values[info.paramName]
	if !ok || len(r) == 0 {
		if info.required {
			return false, info.requiredError
		}
		return false, nil
	}
	return true, p.bindStringSlice(info, expr, r)
}

// NOTE: len(a)>0
func (p *paramInfo) bindStringSlice(info *tagInfo, expr *tagexpr.TagExpr, a []string) error {
	v, err := p.getField(expr, true)
	if err != nil || !v.IsValid() {
		return err
	}

	v = ameda_loc.DereferenceValue(v)

	// we have customized unmarshal defined, we should use it firstly
	if fn, exist := typeUnmarshalFuncs[v.Type()]; exist {
		vv, err := fn(a[0], p.looseZeroMode)
		if err == nil {
			v.Set(vv)
			return nil
		}
		return info.typeError
	}

	switch v.Kind() {
	case reflect.String:
		v.SetString(a[0])
		return nil

	case reflect.Bool:
		var bol bool
		bol, err = strconv.ParseBool(a[0])
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetBool(bol)
			return nil
		}
	case reflect.Float32:
		var f float64
		f, err = strconv.ParseFloat(a[0], 32)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetFloat(f)
			return nil
		}
	case reflect.Float64:
		var f float64
		f, err = strconv.ParseFloat(a[0], 64)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetFloat(f)
			return nil
		}
	case reflect.Int64, reflect.Int:
		var i int64
		i, err = strconv.ParseInt(a[0], 10, 64)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetInt(i)
			return nil
		}
	case reflect.Int32:
		var i int64
		i, err = strconv.ParseInt(a[0], 10, 32)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetInt(i)
			return nil
		}
	case reflect.Int16:
		var i int64
		i, err = strconv.ParseInt(a[0], 10, 16)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetInt(i)
			return nil
		}
	case reflect.Int8:
		var i int64
		i, err = strconv.ParseInt(a[0], 10, 8)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetInt(i)
			return nil
		}
	case reflect.Uint64, reflect.Uint:
		var u uint64
		u, err = strconv.ParseUint(a[0], 10, 64)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetUint(u)
			return nil
		}
	case reflect.Uint32:
		var u uint64
		u, err = strconv.ParseUint(a[0], 10, 32)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetUint(u)
			return nil
		}
	case reflect.Uint16:
		var u uint64
		u, err = strconv.ParseUint(a[0], 10, 16)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetUint(u)
			return nil
		}
	case reflect.Uint8:
		var u uint64
		u, err = strconv.ParseUint(a[0], 10, 8)
		if err == nil || (a[0] == "" && p.looseZeroMode) {
			v.SetUint(u)
			return nil
		}
	case reflect.Slice:
		var ptrDepth int
		t := v.Type().Elem()
		elemKind := t.Kind()
		for elemKind == reflect.Ptr {
			t = t.Elem()
			elemKind = t.Kind()
			ptrDepth++
		}
		val := reflect.New(v.Type()).Elem()
		for _, s := range a {
			var vv reflect.Value
			vv, err = stringToValue(t, s, p.looseZeroMode)
			if err != nil {
				break
			}
			val = reflect.Append(val, ameda_loc.ReferenceValue(vv, ptrDepth))
		}
		if err == nil {
			v.Set(val)
			return nil
		}
		fallthrough
	default:
		// no customized unmarshal defined
		err = unmarshal(ameda_loc.UnsafeStringToBytes(a[0]), v.Addr().Interface())
		if err == nil {
			return nil
		}
	}
	return info.typeError
}

func (p *paramInfo) bindDefaultVal(expr *tagexpr.TagExpr, defaultValue []byte) (bool, error) {
	if defaultValue == nil {
		return false, nil
	}
	v, err := p.getField(expr, true)
	if err != nil || !v.IsValid() {
		return false, err
	}
	return true, jsonpkg.Unmarshal(defaultValue, v.Addr().Interface())
}

// setDefaultVal preprocess the default tags and store the parsed value
func (p *paramInfo) setDefaultVal() error {
	for _, info := range p.tagInfos {
		if info.paramIn != default_val {
			continue
		}

		defaultVal := info.paramName
		st := ameda_loc.DereferenceType(p.structField.Type)
		switch st.Kind() {
		case reflect.String:
			p.defaultVal, _ = jsonpkg.Marshal(defaultVal)
			continue
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
			// escape single quote and double quote, replace single quote with double quote
			defaultVal = strings.Replace(defaultVal, `"`, `\"`, -1)
			defaultVal = strings.Replace(defaultVal, `\'`, specialChar, -1)
			defaultVal = strings.Replace(defaultVal, `'`, `"`, -1)
			defaultVal = strings.Replace(defaultVal, specialChar, `'`, -1)
		}
		p.defaultVal = ameda_loc.UnsafeStringToBytes(defaultVal)
	}
	return nil
}

func stringToValue(elemType reflect.Type, s string, emptyAsZero bool) (v reflect.Value, err error) {
	v = reflect.New(elemType).Elem()

	// we have customized unmarshal defined, we should use it firstly
	if fn, exist := typeUnmarshalFuncs[elemType]; exist {
		vv, err := fn(s, emptyAsZero)
		if err == nil {
			v.Set(vv)
		}
		return v, err
	}

	switch elemType.Kind() {
	case reflect.String:
		v.SetString(s)
	case reflect.Bool:
		var i bool
		i, err = ameda_loc.StringToBool(s, emptyAsZero)
		if err == nil {
			v.SetBool(i)
		}
	case reflect.Float32, reflect.Float64:
		var i float64
		i, err = ameda_loc.StringToFloat64(s, emptyAsZero)
		if err == nil {
			v.SetFloat(i)
		}
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		var i int64
		i, err = ameda_loc.StringToInt64(s, emptyAsZero)
		if err == nil {
			v.SetInt(i)
		}
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		var i uint64
		i, err = ameda_loc.StringToUint64(s, emptyAsZero)
		if err == nil {
			v.SetUint(i)
		}
	default:
		// no customized unmarshal defined
		err = unmarshal(ameda_loc.UnsafeStringToBytes(s), v.Addr().Interface())
	}
	if err != nil {
		return reflect.Value{}, fmt.Errorf("type mismatch, error=%v", err)
	}
	return v, nil
}
