package util

import "testing"

func TestAbs(t *testing.T) {
	want := 2
	got := Abs(-2)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = 9
	got = Abs(9)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = 0
	got = Abs(0)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestMinAll(t *testing.T) {
	nums := []int{5, 1000, 2, 10}
	want := 2
	got := MinAll(nums...)
	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	got = MinAll(2)
	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = 3
	got = MinAll(3, 7)
	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	nums = []int{5, -1000, 2, 10}
	want = -1000
	got = MinAll(nums...)
	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestMaxAll(t *testing.T) {
	nums := []int{5, 1000, 2, 10}
	want := 1000
	got := MaxAll(nums...)
	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	got = MaxAll(2)
	want = 2
	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = 7
	got = MaxAll(3, 7)
	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	nums = []int{5, -1000, 2, 10}
	want = 10
	got = MaxAll(nums...)
	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}
