package ameda_loc

// OneInt8 try to return the first element, otherwise return zero value.
func OneInt8(i []int8) int8 {
	if len(i) > 0 {
		return i[0]
	}
	return 0
}

// Int8sCopy creates a copy of the int8 slice.
func Int8sCopy(i []int8) []int8 {
	b := make([]int8, len(i))
	copy(b, i)
	return b
}

// Int8sToInterfaces converts int8 slice to interface slice.
func Int8sToInterfaces(i []int8) []interface{} {
	r := make([]interface{}, len(i))
	for k, v := range i {
		r[k] = v
	}
	return r
}

// Int8sToStrings converts int8 slice to string slice.
func Int8sToStrings(i []int8) []string {
	r := make([]string, len(i))
	for k, v := range i {
		r[k] = Int8ToString(v)
	}
	return r
}

// Int8sToBools converts int8 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func Int8sToBools(i []int8) []bool {
	r := make([]bool, len(i))
	for k, v := range i {
		r[k] = Int8ToBool(v)
	}
	return r
}

// Int8sToFloat32s converts int8 slice to float32 slice.
func Int8sToFloat32s(i []int8) []float32 {
	r := make([]float32, len(i))
	for k, v := range i {
		r[k] = Int8ToFloat32(v)
	}
	return r
}

// Int8sToFloat64s converts int8 slice to float64 slice.
func Int8sToFloat64s(i []int8) []float64 {
	r := make([]float64, len(i))
	for k, v := range i {
		r[k] = Int8ToFloat64(v)
	}
	return r
}

// Int8sToInts converts int8 slice to int slice.
func Int8sToInts(i []int8) []int {
	r := make([]int, len(i))
	for k, v := range i {
		r[k] = Int8ToInt(v)
	}
	return r
}

// Ints converts int8 slice to int slice.
func Int8sInts(i []int8) []int {
	return Int8sToInts(i)
}

// Int8sToInt16s converts int8 slice to int16 slice.
func Int8sToInt16s(i []int8) []int16 {
	r := make([]int16, len(i))
	for k, v := range i {
		r[k] = Int8ToInt16(v)
	}
	return r
}

// Int8sToInt32s converts int8 slice to int32 slice.
func Int8sToInt32s(i []int8) []int32 {
	r := make([]int32, len(i))
	for k, v := range i {
		r[k] = Int8ToInt32(v)
	}
	return r
}

// Int8sToInt64s converts int8 slice to int64 slice.
func Int8sToInt64s(i []int8) []int64 {
	r := make([]int64, len(i))
	for k, v := range i {
		r[k] = Int8ToInt64(v)
	}
	return r
}

// Int8sToUints converts int8 slice to uint slice.
func Int8sToUints(i []int8) ([]uint, error) {
	var err error
	r := make([]uint, len(i))
	for k, v := range i {
		r[k], err = Int8ToUint(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int8sToUint8s converts int8 slice to uint8 slice.
func Int8sToUint8s(i []int8) ([]uint8, error) {
	var err error
	r := make([]uint8, len(i))
	for k, v := range i {
		r[k], err = Int8ToUint8(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int8sToUint16s converts int8 slice to uint16 slice.
func Int8sToUint16s(i []int8) ([]uint16, error) {
	var err error
	r := make([]uint16, len(i))
	for k, v := range i {
		r[k], err = Int8ToUint16(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int8sToUint32s converts int8 slice to uint32 slice.
func Int8sToUint32s(i []int8) ([]uint32, error) {
	var err error
	r := make([]uint32, len(i))
	for k, v := range i {
		r[k], err = Int8ToUint32(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int8sToUint64s converts int8 slice to uint64 slice.
func Int8sToUint64s(i []int8) ([]uint64, error) {
	var err error
	r := make([]uint64, len(i))
	for k, v := range i {
		r[k], err = Int8ToUint64(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int8sCopyWithin copies part of an slice to another location in the current slice.
// @target
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func Int8sCopyWithin(i []int8, target, start int, end ...int) {
	target = fixIndex(len(i), target, true)
	if target == len(i) {
		return
	}
	sub := Int8sSlice(i, start, end...)
	for k, v := range sub {
		i[target+k] = v
	}
}

// Int8sEvery tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func Int8sEvery(i []int8, fn func(i []int8, k int, v int8) bool) bool {
	for k, v := range i {
		if !fn(i, k, v) {
			return false
		}
	}
	return true
}

// Int8sFill changes all elements in the current slice to a value, from a start index to an end index.
// @value
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func Int8sFill(i []int8, value int8, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(i), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		i[k] = value
	}
}

// Int8sFilter creates a new slice with all elements that pass the test implemented by the provided function.
func Int8sFilter(i []int8, fn func(i []int8, k int, v int8) bool) []int8 {
	ret := make([]int8, 0)
	for k, v := range i {
		if fn(i, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Int8sFind returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func Int8sFind(i []int8, fn func(i []int8, k int, v int8) bool) (k int, v int8) {
	for k, v := range i {
		if fn(i, k, v) {
			return k, v
		}
	}
	return -1, 0
}

// Int8sIncludes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func Int8sIncludes(i []int8, valueToFind int8, fromIndex ...int) bool {
	return Int8sIndexOf(i, valueToFind, fromIndex...) > -1
}

// Int8sIndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func Int8sIndexOf(i []int8, searchElement int8, fromIndex ...int) int {
	idx := getFromIndex(len(i), fromIndex...)
	for k, v := range i[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// Int8sLastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func Int8sLastIndexOf(i []int8, searchElement int8, fromIndex ...int) int {
	idx := getFromIndex(len(i), fromIndex...)
	for k := len(i) - 1; k >= idx; k-- {
		if searchElement == i[k] {
			return k
		}
	}
	return -1
}

// Int8sMap creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func Int8sMap(i []int8, fn func(i []int8, k int, v int8) int8) []int8 {
	ret := make([]int8, len(i))
	for k, v := range i {
		ret[k] = fn(i, k, v)
	}
	return ret
}

// Int8sPop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func Int8sPop(i *[]int8) (int8, bool) {
	a := *i
	if len(a) == 0 {
		return 0, false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*i = a[:len(a):len(a)]
	return last, true
}

// Int8sPush adds one or more elements to the end of an slice and returns the new length of the slice.
func Int8sPush(i *[]int8, element ...int8) int {
	*i = append(*i, element...)
	return len(*i)
}

// Int8sPushDistinct adds one or more new elements that do not exist in the current slice at the end.
func Int8sPushDistinct(i []int8, element ...int8) []int8 {
L:
	for _, v := range element {
		for _, vv := range i {
			if vv == v {
				continue L
			}
		}
		i = append(i, v)
	}
	return i
}

// Int8sReduce executes a reducer function (that you provide) on each element of the slice,
// resulting in a single output value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func Int8sReduce(i []int8,
	fn func(i []int8, k int, v, accumulator int8) int8, initialValue ...int8,
) int8 {
	if len(i) == 0 {
		return 0
	}
	start := 0
	acc := i[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for k := start; k < len(i); k++ {
		acc = fn(i, k, i[k], acc)
	}
	return acc
}

// Int8sReduceRight applies a function against an accumulator and each value of the slice (from right-to-left)
// to reduce it to a single value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func Int8sReduceRight(i []int8,
	fn func(i []int8, k int, v, accumulator int8) int8, initialValue ...int8,
) int8 {
	if len(i) == 0 {
		return 0
	}
	end := len(i) - 1
	acc := i[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for k := end; k >= 0; k-- {
		acc = fn(i, k, i[k], acc)
	}
	return acc
}

// Int8sReverse reverses an slice in place.
func Int8sReverse(i []int8) {
	first := 0
	last := len(i) - 1
	for first < last {
		i[first], i[last] = i[last], i[first]
		first++
		last--
	}
}

// Int8sShift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func Int8sShift(i *[]int8) (int8, bool) {
	a := *i
	if len(a) == 0 {
		return 0, false
	}
	first := a[0]
	a = a[1:]
	*i = a[:len(a):len(a)]
	return first, true
}

// Int8sSlice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func Int8sSlice(i []int8, begin int, end ...int) []int8 {
	fixedStart, fixedEnd, ok := fixRange(len(i), begin, end...)
	if !ok {
		return []int8{}
	}
	return Int8sCopy(i[fixedStart:fixedEnd])
}

// Int8sSome tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func Int8sSome(i []int8, fn func(i []int8, k int, v int8) bool) bool {
	for k, v := range i {
		if fn(i, k, v) {
			return true
		}
	}
	return false
}

// Int8sSplice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func Int8sSplice(i *[]int8, start, deleteCount int, items ...int8) {
	a := *i
	if deleteCount < 0 {
		deleteCount = 0
	}
	start, end, _ := fixRange(len(a), start, start+1+deleteCount)
	deleteCount = end - start - 1
	for k := 0; k < len(items); k++ {
		if deleteCount > 0 {
			// replace
			a[start] = items[k]
			deleteCount--
			start++
		} else {
			// insert
			lastSlice := Int8sCopy(a[start:])
			items = items[k:]
			a = append(a[:start], items...)
			a = append(a[:start+len(items)], lastSlice...)
			*i = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*i = a[:len(a):len(a)]
}

// Int8sUnshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func Int8sUnshift(i *[]int8, element ...int8) int {
	*i = append(element, *i...)
	return len(*i)
}

// Int8sUnshiftDistinct adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func Int8sUnshiftDistinct(i *[]int8, element ...int8) int {
	a := *i
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[int8]bool, len(element))
	r := make([]int8, 0, len(a)+len(element))
L:
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		r = append(r, v)
	}
	r = append(r, a...)
	*i = r[:len(r):len(r)]
	return len(r)
}

// Int8sRemoveFirst removes the first matched elements from the slice,
// and returns the new length of the slice.
func Int8sRemoveFirst(p *[]int8, elements ...int8) int {
	a := *p
	m := make(map[interface{}]struct{}, len(elements))
	for _, element := range elements {
		if _, ok := m[element]; ok {
			continue
		}
		m[element] = struct{}{}
		for k, v := range a {
			if v == element {
				a = append(a[:k], a[k+1:]...)
				break
			}
		}
	}
	n := len(a)
	*p = a[:n:n]
	return n
}

// Int8sRemoveEvery removes all the elements from the slice,
// and returns the new length of the slice.
func Int8sRemoveEvery(p *[]int8, elements ...int8) int {
	a := *p
	m := make(map[interface{}]struct{}, len(elements))
	for _, element := range elements {
		if _, ok := m[element]; ok {
			continue
		}
		m[element] = struct{}{}
		for i := 0; i < len(a); i++ {
			if a[i] == element {
				a = append(a[:i], a[i+1:]...)
				i--
			}
		}
	}
	n := len(a)
	*p = a[:n:n]
	return n
}

// Int8sConcat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func Int8sConcat(i ...[]int8) []int8 {
	var totalLen int
	for _, v := range i {
		totalLen += len(v)
	}
	ret := make([]int8, totalLen)
	dst := ret
	for _, v := range i {
		n := copy(dst, v)
		dst = dst[n:]
	}
	return ret
}

// Int8sIntersect calculates intersection of two or more slices,
// and returns the count of each element.
func Int8sIntersect(i ...[]int8) (intersectCount map[int8]int) {
	if len(i) == 0 {
		return nil
	}
	for _, v := range i {
		if len(v) == 0 {
			return nil
		}
	}
	counts := make([]map[int8]int, len(i))
	for k, v := range i {
		counts[k] = int8sDistinct(v, nil)
	}
	intersectCount = counts[0]
L:
	for k, v := range intersectCount {
		for _, c := range counts[1:] {
			v2 := c[k]
			if v2 == 0 {
				delete(intersectCount, k)
				continue L
			}
			if v > v2 {
				v = v2
			}
		}
		intersectCount[k] = v
	}
	return intersectCount
}

// Int8sDistinct calculates the count of each different element,
// and only saves these different elements in place if changeSlice is true.
func Int8sDistinct(i *[]int8, changeSlice bool) (distinctCount map[int8]int) {
	if !changeSlice {
		return int8sDistinct(*i, nil)
	}
	a := (*i)[:0]
	distinctCount = int8sDistinct(*i, &a)
	n := len(distinctCount)
	*i = a[:n:n]
	return distinctCount
}

func int8sDistinct(src []int8, dst *[]int8) map[int8]int {
	m := make(map[int8]int, len(src))
	if dst == nil {
		for _, v := range src {
			n := m[v]
			m[v] = n + 1
		}
	} else {
		a := *dst
		for _, v := range src {
			n := m[v]
			m[v] = n + 1
			if n == 0 {
				a = append(a, v)
			}
		}
		*dst = a
	}
	return m
}

// Int8SetUnion calculates between multiple collections: set1 ∪ set2 ∪ others...
// This method does not change the existing slices, but instead returns a new slice.
func Int8SetUnion(set1, set2 []int8, others ...[]int8) []int8 {
	m := make(map[int8]struct{}, len(set1)+len(set2))
	r := make([]int8, 0, len(m))
	for _, set := range append([][]int8{set1, set2}, others...) {
		for _, v := range set {
			_, ok := m[v]
			if ok {
				continue
			}
			r = append(r, v)
			m[v] = struct{}{}
		}
	}
	return r
}

// Int8SetIntersect calculates between multiple collections: set1 ∩ set2 ∩ others...
// This method does not change the existing slices, but instead returns a new slice.
func Int8SetIntersect(set1, set2 []int8, others ...[]int8) []int8 {
	sets := append([][]int8{set2}, others...)
	setsCount := make([]map[int8]int, len(sets))
	for k, v := range sets {
		setsCount[k] = int8sDistinct(v, nil)
	}
	m := make(map[int8]struct{}, len(set1))
	r := make([]int8, 0, len(m))
L:
	for _, v := range set1 {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		for _, m2 := range setsCount {
			if m2[v] == 0 {
				continue L
			}
		}
		r = append(r, v)
	}
	return r
}

// Int8SetDifference calculates between multiple collections: set1 - set2 - others...
// This method does not change the existing slices, but instead returns a new slice.
func Int8SetDifference(set1, set2 []int8, others ...[]int8) []int8 {
	m := make(map[int8]struct{}, len(set1))
	r := make([]int8, 0, len(set1))
	sets := append([][]int8{set2}, others...)
	for _, v := range sets {
		inter := Int8SetIntersect(set1, v)
		for _, v := range inter {
			m[v] = struct{}{}
		}
	}
	for _, v := range set1 {
		if _, ok := m[v]; !ok {
			r = append(r, v)
			m[v] = struct{}{}
		}
	}
	return r
}
