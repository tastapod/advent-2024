package ints

import "strconv"

func AbsInt(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func ToInt(s string) (result int) {
	result, _ = strconv.Atoi(s)
	return
}
