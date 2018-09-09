package utils

import "strconv"

// EqualsStringSlices checks if two slices of strings are equals or not.
func EqualsStringSlices(a []string, b []string) bool {
	if len(a) != len(b) || cap(a) != cap(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// StringToInt converts a string to the corresponding int or panic if it's not a valid int string.
func StringToInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
