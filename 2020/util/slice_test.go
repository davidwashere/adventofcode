package util

import "testing"

func TestRemoveIndexFromStringSlice(t *testing.T) {
	vf := func(got, want interface{}) {
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}

	slice := []string{"a", "b", "c", "d"}
	res := RemoveIndexFromStringSlice(slice, 1) // remove 'b'

	vf(len(res), 3)
	vf(res[0], "a")
	vf(res[1], "c")
	vf(res[2], "d")

	res = RemoveIndexFromStringSlice(res, 2) // remove 'd'

	vf(len(res), 2)
	vf(res[0], "a")
	vf(res[1], "c")

	res = RemoveIndexFromStringSlice(res, 0) // remove 'a'

	vf(len(res), 1)
	vf(res[0], "c")

	res = RemoveIndexFromStringSlice(res, 0) // remove 'c'
	vf(len(res), 0)
}

func TestRemoveIndexFromIntSlice(t *testing.T) {
	vf := func(got, want interface{}) {
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}

	slice := []int{1, 2, 3, 4}
	res := RemoveIndexFromIntSlice(slice, 1) // remove '2'

	vf(len(res), 3)
	vf(res[0], 1)
	vf(res[1], 3)
	vf(res[2], 4)

	res = RemoveIndexFromIntSlice(res, 2) // remove '4'

	vf(len(res), 2)
	vf(res[0], 1)
	vf(res[1], 3)

	res = RemoveIndexFromIntSlice(res, 0) // remove '1'

	vf(len(res), 1)
	vf(res[0], 3)

	res = RemoveIndexFromIntSlice(res, 0) // remove '3'
	vf(len(res), 0)
}

func TestRemoveStringFromSlice(t *testing.T) {
	vf := func(got, want interface{}) {
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}

	slice := []string{"a", "b", "c", "d"}
	res := RemoveStringFromSlice(slice, "b")

	vf(len(res), 3)
	vf(res[0], "a")
	vf(res[1], "c")
	vf(res[2], "d")

	res = RemoveStringFromSlice(res, "d")

	vf(len(res), 2)
	vf(res[0], "a")
	vf(res[1], "c")

	res = RemoveStringFromSlice(res, "a")

	vf(len(res), 1)
	vf(res[0], "c")

	res = RemoveStringFromSlice(res, "c")
	vf(len(res), 0)

}

func TestRemoveIntFromSlice(t *testing.T) {
	vf := func(got, want interface{}) {
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}

	slice := []int{1, 2, 3, 4}
	res := RemoveIntFromSlice(slice, 2)

	vf(len(res), 3)
	vf(res[0], 1)
	vf(res[1], 3)
	vf(res[2], 4)

	res = RemoveIntFromSlice(res, 4)

	vf(len(res), 2)
	vf(res[0], 1)
	vf(res[1], 3)

	res = RemoveIntFromSlice(res, 1)

	vf(len(res), 1)
	vf(res[0], 3)

	res = RemoveIntFromSlice(res, 3)
	vf(len(res), 0)
}

func TestIsStringIn(t *testing.T) {
	vf := func(got, want interface{}) {
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}

	slice := []string{"a", "b", "c", "d"}

	vf(IsStringIn(slice, "a"), true)
	vf(IsStringIn(slice, "x"), false)
}

func TestIsIntIn(t *testing.T) {
	vf := func(got, want interface{}) {
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}

	slice := []int{1, 2, 3, 4}

	vf(IsIntIn(slice, 3), true)
	vf(IsIntIn(slice, 0), false)
}
