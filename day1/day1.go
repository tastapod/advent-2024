package day1

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

func ParseInput(input string) (l, r []int) {
	rows := strings.Split(input, "\n")
	l = make([]int, len(rows))
	r = make([]int, len(rows))
	for i, row := range rows {
		parts := strings.Fields(row)
		l[i], _ = strconv.Atoi(parts[0])
		r[i], _ = strconv.Atoi(parts[1])
	}
	return
}

func SumDeltas(l, r []int) int {
	lSorted := sorted(l)
	rSorted := sorted(r)

	total := 0
	for i := 0; i < len(lSorted); i++ {
		total += abs(lSorted[i] - rSorted[i])
	}
	return total
}

func SimilarityScore(l, r []int) (result int) {
	counts := BuildHistogram(r)

	for _, num := range l {
		result += num * counts[num]
	}
	return
}

func BuildHistogram(nums []int) (result map[int]int) {
	result = make(map[int]int)
	numsSorted := sorted(nums)

	count := 1
	current := numsSorted[0]
	for i := 1; i < len(numsSorted); i++ {
		if this := numsSorted[i]; this == current {
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
