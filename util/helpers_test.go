package util

import "testing"

var vf = func(t *testing.T, got, want any) {
	if got != want {
		t.Helper()
		t.Errorf("Got %v want %v", got, want)
	}
}
