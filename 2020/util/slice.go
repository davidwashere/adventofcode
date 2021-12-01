package util

// RemoveIndexFromStringSlice removes the string found at `index`
func RemoveIndexFromStringSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveIndexFromIntSlice removes the int found at `index`
func RemoveIndexFromIntSlice(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveStringFromSlice removes `val` from slice if found, if not found returns
// the slice as is
func RemoveStringFromSlice(slice []string, val string) []string {
	for i, s := range slice {
		if s == val {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

// RemoveIntFromSlice removes `val` from slice if found, if not found returns
// the slice as is
func RemoveIntFromSlice(slice []int, val int) []int {
	for i, s := range slice {
		if s == val {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

// IsStringIn will return true if val is found in slice, false otherwise
func IsStringIn(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}

	return false
}

// IsIntIn will return true if val is found in slice, false otherwise
func IsIntIn(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}

	return false
}
