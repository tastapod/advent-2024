package parsing

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadDay(day int) string {
	dayFile := fmt.Sprintf("day%d/input.txt", day)

	content, err := os.ReadFile(dayFile)
	if err != nil {
		panic(fmt.Sprintf("%s: error %s", dayFile, err))
	}
	return strings.TrimSpace(string(content))
}

func ReadAndSplitDay(day int) []string {
	return strings.Split(ReadDay(day), "\n")
}

func Parts(input string) []string {
	return PartsWithSep(input, "\n")
}

func PartsWithSep(input, sep string) []string {
	return strings.Split(strings.TrimSpace(input), sep)
}

func Ints(input string) (result []int) {

	fields := strings.Fields(input)
	result = make([]int, len(fields))

	for i, val := range fields {
		result[i], _ = strconv.Atoi(val)
	}
	return
}

func IntsWithSep(input, sep string) (result []int) {
	parts := PartsWithSep(input, sep)
	result = make([]int, len(parts))

	for i, part := range parts {
		result[i] = Int(part)
	}
	return
}

func Int(s string) (result int) {
	result, _ = strconv.Atoi(s)
	return
}
