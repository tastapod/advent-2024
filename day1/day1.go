package day1

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type ListPair struct {
	L, R      []int
	Histogram map[int]int
}

func NewListPair(input string) (pair *ListPair) {
	rows := strings.Split(input, "\n")
	pair = &ListPair{
		L:         make([]int, len(rows)),
		R:         make([]int, len(rows)),
		Histogram: make(map[int]int),
	}
	for i, row := range rows {
		parts := strings.Fields(row)
		pair.L[i], _ = strconv.Atoi(parts[0])
		pair.R[i], _ = strconv.Atoi(parts[1])
	}
	slices.Sort(pair.L)
	slices.Sort(pair.R)

	// Build histogram
	count := 1
	current := pair.R[0]
	for _, this := range pair.R[1:] {
		if this == current {
			count++
		} else {
			pair.Histogram[current] = count
			count = 1
			current = this
		}
	}
	pair.Histogram[current] = count
	return
}

func (pair *ListPair) SumDeltas() int {
	lSorted := sorted(pair.L)
	rSorted := sorted(pair.R)

	total := 0
	for i := 0; i < len(lSorted); i++ {
		total += abs(lSorted[i] - rSorted[i])
	}
	return total
}

func (pair *ListPair) SimilarityScore() (result int) {
	for _, num := range pair.L {
		result += num * pair.Histogram[num]
	}
	return
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func sorted[T cmp.Ordered](l []T) []T {
	lSorted := l[:]
	slices.Sort(lSorted)
	return lSorted
}
