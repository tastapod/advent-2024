package parsing

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadDay reads a file called `input.txt` from the given day's source directory
func ReadDay(day int) string {
	return TrimFile(fmt.Sprintf("day%d/input.txt", day))
}

func TrimFile(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("%s: error %s", file, err))
	}
	return strings.TrimSpace(string(content))
}

func FileLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("%s: error %s", file, err))
	}
	return Lines(string(content))
}

// ReadAndSplitDay reads a file called `input.txt` and splits it on newlines
func ReadAndSplitDay(day int) []string {
	return strings.Split(ReadDay(day), "\n")
}

// Lines trims a string and splits it at newlines
func Lines(input string) []string {
	return Parts(input, "\n")
}

// Parts trims a string and splits it at the separator
func Parts(input, sep string) []string {
	return strings.Split(strings.TrimSpace(input), sep)
}

// Int is a convenience method to convert a string to int ignoring errors
func Int(value string) (result int) {
	result, _ = strconv.Atoi(value)
	return
}

// Ints splits a string into fields and converts them to ints
func Ints(input string) (result []int) {

	fields := strings.Fields(input)
	result = make([]int, len(fields))

	for i, val := range fields {
		result[i] = Int(val)
	}
	return
}

// Int64 is a convenience method to convert a string to int64 ignoring errors
func Int64(value string) int64 {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("%s: error %s", value, err))
	}
	return result
}

// Int64s splits a string into fields and converts them to int64
func Int64s(input string) (result []int64) {

	fields := strings.Fields(input)
	result = make([]int64, len(fields))

	for i, val := range fields {
		result[i] = Int64(val)
	}
	return
}
