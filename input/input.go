package input

import (
	"fmt"
	"os"
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
