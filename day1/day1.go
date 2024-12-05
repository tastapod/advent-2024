package day1

import (
	"github.com/tastapod/advent-2024/internal/ints"
	"slices"
	"strconv"
	"strings"
)

type ListPair struct {
	L, R []int
}

func NewListPair(input string) (pair *ListPair) {
	rows := strings.Split(input, "\n")
	pair = &ListPair{
		L: make([]int, len(rows)),
		R: make([]int, len(rows)),
	}
	for i, row := range rows {
		parts := strings.Fields(row)
		pair.L[i], _ = strconv.Atoi(parts[0])
		pair.R[i], _ = strconv.Atoi(parts[1])
	}
	slices.Sort(pair.L)
	slices.Sort(pair.R)
	return
}

func (pair *ListPair) SumDeltas() int {
	total := 0
	for i := 0; i < len(pair.L); i++ {
		total += ints.Abs(pair.L[i] - pair.R[i])
	}
	return total
}

func (pair *ListPair) SimilarityScore() (result int) {
	lHist := RunLengthHistogram(pair.L)
	rHist := RunLengthHistogram(pair.R)

	// This is just a sum of products
	for value, lCount := range lHist {
		result += value * lCount * rHist[value]
	}
	return
}

// RunLengthHistogram builds a histogram of run-length counts
// So [1, 1, 1, 4, 4] -> {1: 3, 4: 2}
func RunLengthHistogram[T comparable](sortedValues []T) (result map[T]int) {
	result = make(map[T]int)
	count := 1
	current := sortedValues[0]
	for _, this := range sortedValues[1:] {
		if this == current {
			count++
		} else {
			result[current] = count
			count = 1
			current = this
		}
	}
	result[current] = count
	return
}
