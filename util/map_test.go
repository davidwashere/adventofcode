package util

import (
	"reflect"
	"testing"
)

func TestSortMapKeys(t *testing.T) {
	m := map[string]int{}

	m["1"] = 0
	m["10"] = 0
	m["2"] = 0

	want := []string{"1", "2", "10"}
	// want := []string{"1", "10", "2"}
	got := SortMapKeysInt(m)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSortMapKeysLower(t *testing.T) {
	m := map[string]int{}

	m["A"] = 0
	m["b"] = 0
	m["C"] = 0

	want := []string{"A", "b", "C"}
	// want := []string{"1", "10", "2"}
	got := SortMapKeysLower(m)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
