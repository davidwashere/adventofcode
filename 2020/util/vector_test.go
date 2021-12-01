package util

import (
	"testing"
)

func TestRotate(t *testing.T) {
	v := NewVector(2, 1, 1)

	v.Rotate(90)
	wantX, wantY := -1, 2
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}

	v.Rotate(180)
	wantX, wantY = 1, -2
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}

	v.Rotate(90)
	wantX, wantY = 2, 1
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}

	v.Rotate(-90)
	wantX, wantY = 1, -2
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}

	v = NewVector(0, 1, 1)
	v.Rotate(-45)
	wantX, wantY = 1, 1
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}

	v.Rotate(90)
	wantX, wantY = -1, 1
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}

	v.Rotate(360)
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}

	v.Rotate(720)
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}

	v = NewVector(0, 1, 1)
	v.Rotate(540)
	wantX, wantY = 0, -1
	if wantX != v.X || wantY != v.Y {
		t.Errorf("Got %v, %v want %v, %v", v.X, v.Y, wantX, wantY)
	}
}
