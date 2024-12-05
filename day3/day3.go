package day3

import (
	"github.com/tastapod/advent-2024/internal/parsing"
	"regexp"
)

const MulRE = `mul\((\d+),(\d+)\)`
const DoRE = `do\(\)`
const DontRE = `don't\(\)`

func FindMuls(input string) [][]string {
	return regexp.MustCompile(MulRE).FindAllStringSubmatch(input, -1)
}

func Multiply(match []string) int {
	return parsing.Int(match[1]) * parsing.Int(match[2])
}

func SumMuls(muls [][]string) (result int) {
	for _, mul := range muls {
		result += Multiply(mul)
	}
	return
}

func SumEnabledMuls(input string) (result int) {
	re := regexp.MustCompile(MulRE + "|" + DoRE + "|" + DontRE)

	enabled := true
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		switch match[0] {
		case `do()`:
			enabled = true
		case `don't()`:
			enabled = false
		default:
			if enabled {
				result += Multiply(match)
			}
		}
	}
	return
}
