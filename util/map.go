package util

import (
	"reflect"
	"sort"
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
