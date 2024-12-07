package parsing

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadDay reads a file called `input.txt` from the given day's source directory
func ReadDay(day int) string {
	dayFile := fmt.Sprintf("day%d/input.txt", day)

	content, err := os.ReadFile(dayFile)
	if err != nil {
		panic(fmt.Sprintf("%s: error %s", dayFile, err))
	}
	return strings.TrimSpace(string(content))
}

// ReadAndSplitDay reads a file called `input.txt` and splits it on newlines
func ReadAndSplitDay(day int) []string {
	return strings.Split(ReadDay(day), "\n")
}

// Parts trims a string and splits it at newlines
func Parts(input string) []string {
	return PartsWithSep(input, "\n")
}

// PartsWithSep trims a string and splits it at the separator
func PartsWithSep(input, sep string) []string {
	return strings.Split(strings.TrimSpace(input), sep)
}

// Ints splits a string into fields and converts them to ints
func Ints(input string) (result []int) {

	fields := strings.Fields(input)
	result = make([]int, len(fields))

	for i, val := range fields {
		result[i], _ = strconv.Atoi(val)
	}
	return
}

// IntsWithSep splits a string on a separator and converts the parts to ints
func IntsWithSep(input, sep string) (result []int) {
	parts := PartsWithSep(input, sep)
	result = make([]int, len(parts))

	for i, part := range parts {
		result[i] = Int(part)
	}
	return
}

// Int is a convenience method to convert a string to int ignoring errors
func Int(s string) (result int) {
	result, _ = strconv.Atoi(s)
	return
}
