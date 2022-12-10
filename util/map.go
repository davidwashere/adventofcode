package util

import (
	"reflect"
	"sort"
	"strconv"
)

// SortMapKeys accepts a map who's keys are strings, map[string]... and
// returns a slice of the keys sorted
func SortMapKeys(in interface{}) []string {
	v := reflect.ValueOf(in)

	keys := []string{}

	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			keys = append(keys, key.Interface().(string))
		}
	}

	sort.Strings(keys)

	return keys
}

// SortMapKeys accepts a map who's keys are strings, map[string]... and
// returns a slice of the keys sorted based on interpreting the keys as ints
func SortMapKeysInt(in interface{}) []string {
	v := reflect.ValueOf(in)

	keys := []string{}

	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			keys = append(keys, key.Interface().(string))
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		a, _ := strconv.Atoi(keys[i])
		b, _ := strconv.Atoi(keys[j])

		return a < b
	})

	return keys
}
