package util

import "reflect"

// PermsString same as `Perms` but internally converts data to `string`
func PermsString(inSlice interface{}, repeat bool, f func([]string)) {
	PermsStringOfLen(inSlice, -1, repeat, f)
}

// PermsStringOfLen same as `PermsOfLen` but internally converts data to `string`
func PermsStringOfLen(inSlice interface{}, maxlen int, repeat bool, f func([]string)) {
	PermsOfLen(inSlice, maxlen, repeat, func(perm []interface{}) {
		conv := make([]string, len(perm))
		for i := 0; i < len(perm); i++ {
			conv[i] = perm[i].(string)
		}
		f(conv)
	})
}

// PermsInt same as `Perms` but internally converts data to `int`
func PermsInt(inSlice interface{}, repeat bool, f func([]int)) {
	PermsIntOfLen(inSlice, -1, repeat, f)
}

// PermsIntOfLen same as `PermsOfLen` but internally converts data to `int`
func PermsIntOfLen(inSlice interface{}, maxlen int, repeat bool, f func([]int)) {
	PermsOfLen(inSlice, maxlen, repeat, func(perm []interface{}) {
		conv := make([]int, len(perm))
		for i := 0; i < len(perm); i++ {
			conv[i] = perm[i].(int)
		}
		f(conv)
	})
}

// Perms will call `f` for every permutation of the items in `inSlice`
func Perms(inSlice interface{}, repeat bool, f func([]interface{})) {
	PermsOfLen(inSlice, -1, repeat, f)
}

// PermsOfLen same as `Perms` but will create permutations of a set length
func PermsOfLen(inSlice interface{}, maxlen int, repeat bool, f func([]interface{})) {
	slice, success := convToType(inSlice, reflect.Slice)

	if !success {
		panic("Slice conversion no worky")
	}

	var convertedData []interface{}
	length := slice.Len()
	convertedData = make([]interface{}, length)
	for i := 0; i < length; i++ {
		convertedData[i] = slice.Index(i).Interface()
	}

	if maxlen < 0 {
		maxlen = length
	}

	if repeat {
		permsRepeat(convertedData, nil, maxlen, f)
	} else {
		if maxlen > length {
			maxlen = length
		}
		perms(convertedData, nil, maxlen, f)

	}
}

// perms calls f for every permutation up to len l without repeats
func perms(data []interface{}, ans []interface{}, l int, f func([]interface{})) {
	if ans != nil && len(ans) == l {
		f(ans)
		return
	}

	for i := 0; i < len(data); i++ {
		item := data[i]

		var rod []interface{}
		rod = append(rod, data[0:i]...)
		rod = append(rod, data[i+1:]...)

		var newans []interface{}
		newans = append(newans, ans...)
		newans = append(newans, item)

		perms(rod, newans, l, f)
	}
}

// permsRepeat calls f for every permutation up to len l with repeats
func permsRepeat(data []interface{}, ans []interface{}, l int, f func([]interface{})) {
	for i := 0; i < len(data); i++ {
		var newans []interface{}
		newans = append(newans, ans...)
		newans = append(newans, data[i])

		if len(newans) == l {
			f(newans)
		} else {
			permsRepeat(data, newans, l, f)
		}
	}
}

// PermsChar same as `Perms` but looks at characters inside a string (instead of expecting a slice)
// set repeat to true if chars can repeat, false otherwise
func PermsChar(str string, repeat bool, f func(string)) {
	PermsCharOfLen(str, -1, repeat, f)
}

// PermsCharOfLen same as `PermsChar` but will limit permutations to the specific length
// set repeat to true if chars can repeat, false otherwise
func PermsCharOfLen(str string, maxlen int, repeat bool, f func(string)) {
	if maxlen < 0 {
		maxlen = len(str)
	}

	if repeat {
		permsCharRepeat(str, "", maxlen, f)
	} else {
		if maxlen > len(str) {
			maxlen = len(str)
		}
		permsChar(str, "", maxlen, f)
	}
}

// permsChar calls f for every permutation up to len l without repeats
func permsChar(str string, ans string, l int, f func(string)) {
	if len(ans) == l {
		f(ans)
		return
	}

	for i := 0; i < len(str); i++ {
		chr := string(str[i])
		ros := str[0:i] + str[i+1:]
		permsChar(ros, ans+chr, l, f)
	}
}

// permsChar calls f for every permutation up to len l with repeats
func permsCharRepeat(str string, ans string, l int, f func(string)) {
	for i := 0; i < len(str); i++ {
		newans := ans + string(str[i])

		if len(newans) == l {
			f(newans)
		} else {
			permsCharRepeat(str, newans, l, f)
		}
	}
}
