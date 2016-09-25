package service

import "strconv"

func equalsStringSlices(a []string, b []string) bool {
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

func stringToInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
