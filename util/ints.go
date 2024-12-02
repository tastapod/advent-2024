package util

import "strconv"

func ToInt(num string) int {
	value, _ := strconv.Atoi(num)
	return value
}

func AbsInt(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
