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

func ParseInts(input string) (result []int) {

	fields := strings.Fields(input)
	result = make([]int, len(fields))

	for i, val := range fields {
		result[i], _ = strconv.Atoi(val)
	}
	return
}
