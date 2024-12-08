package parsing

import (
	"fmt"
	"math/big"
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

// IntsWithSep splits a string on a separator and converts the parts to ints
func IntsWithSep(input, sep string) (result []int) {
	parts := Parts(input, sep)
	result = make([]int, len(parts))

	for i, part := range parts {
		result[i] = Int(part)
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

func BigInt(value string) *big.Int {
	result, ok := big.NewInt(0).SetString(value, 10)
	if !ok {
		panic(fmt.Sprintf("error %s", value))
	}
	return result
}

func BigInts(values string) (result []*big.Int) {
	fields := strings.Fields(values)
	result = make([]*big.Int, len(fields))
	for i, field := range fields {
		result[i] = BigInt(field)
	}
	return
}
