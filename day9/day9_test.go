package day9_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tastapod/advent-2024/day9"
	"github.com/tastapod/advent-2024/internal/parsing"
	"testing"
)

/*
Build a run-length map that consumes from the right (this is like filling trades!)

- files: [id: 0, start: 0, len: 1], [id: 1, start: 3, len: 3], [id: 2, start:10, len 5]
- spaces: [start: 1, len: 2], [start:6, len 4]

or: linked list of file-or-space

[file0 -> space -> file1 -> space -> ...]
Then to fill from the right, start with right-most (last) file
- Fill spaces
	- exact fill: space -> file (file done)
	- partial fill: space -> file + new space (file done)
	- over fill: space -> part file; file -> remainder file

*/

func TestMakesDiskMap(t *testing.T) {
	// given
	line := "12345"

	// when
	diskMap := day9.NewDiskMap(line)

	// then
	assert.Equal(t, 5, len(diskMap.Entries))

	expected := []day9.DiskEntry{
		{0, 0, 1},
		{-1, 1, 2},
		{1, 3, 3},
		{-1, 6, 4},
		{2, 10, 5},
	}

	assert.Equal(t, expected, diskMap.Entries)
}

func TestDefragsMatchingLength(t *testing.T) {
	// given
	dm := day9.NewDiskMap("133") // 1-file, 3-space 3-file
	assert.Equal(t, []day9.DiskEntry{
		{0, 0, 1},
		{-1, 1, 3},
		{1, 4, 3},
	}, dm.Entries)

	// when
	dm.DefragLastFile()

	// then
	assert.Equal(t, []day9.DiskEntry{
		{0, 0, 1},
		{1, 1, 3},
		{-1, 4, 3},
	}, dm.Entries)
}

func TestDefragsLongFile(t *testing.T) {
	// given
	dm := day9.NewDiskMap("135") // 1-file, 3-space 5-file
	assert.Equal(t, []day9.DiskEntry{
		{0, 0, 1},
		{-1, 1, 3},
		{1, 4, 5},
	}, dm.Entries)

	// when
	dm.DefragLastFile()

	// then
	assert.Equal(t, []day9.DiskEntry{
		{0, 0, 1},
		{1, 1, 3},
		{1, 4, 2},
		{-1, 7, 3},
	}, dm.Entries)
}

func TestDefragsShortFile(t *testing.T) {
	// given
	dm := day9.NewDiskMap("152") // 1-file, 5-space 3-file
	assert.Equal(t, []day9.DiskEntry{
		{0, 0, 1},
		{-1, 1, 5},
		{1, 6, 2},
	}, dm.Entries)

	// when
	dm.DefragLastFile()

	// then
	assert.Equal(t, []day9.DiskEntry{
		{0, 0, 1},
		{1, 1, 2},
		{-1, 3, 3},
		{-1, 6, 2},
	}, dm.Entries)
}

const Sample = "2333133121414131402"

func TestDefragsSampleDiskMap(t *testing.T) {
	// given
	dm := day9.NewDiskMap(Sample)

	// when
	dm.DefragWholeDisk()

	// then
	assert.Equal(t, "0099811188827773336446555566..............", dm.String())
}

func TestCalculatesChecksumForEntry(t *testing.T) {
	// given
	type TestCase struct {
		day9.DiskEntry
		Expected int
	}

	tcs := []TestCase{
		{day9.DiskEntry{Id: 0, Start: 0, Length: 2}, 0},
		{day9.DiskEntry{Id: 9, Start: 2, Length: 2}, 45},
		{day9.DiskEntry{Id: -1, Start: 2, Length: 2}, 0},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.Expected, tc.DiskEntry.Checksum(), tc.DiskEntry)
	}
}

func TestCalculatesChecksumForDisk(t *testing.T) {
	// given
	dm := day9.NewDiskMap(Sample)
	dm.DefragWholeDisk()

	// then
	assert.Equal(t, 1928, dm.Checksum())
}

func TestCalculatesChecksumForPart1(t *testing.T) {
	// given
	dm := day9.NewDiskMap(parsing.TrimFile("input.txt"))
	dm.DefragWholeDisk()

	// then
	assert.Equal(t, 6307275788409, dm.Checksum())
}

func TestCompactsSpace(t *testing.T) {
	// given
	dm := day9.DiskMap{
		Entries: []day9.DiskEntry{
			{Id: 0, Start: 0, Length: 4},
			{Id: -1, Start: 4, Length: 3},
			{Id: -1, Start: 7, Length: 2},
			{Id: 1, Start: 9, Length: 1},
		}}

	// when
	dm.Compact()

	// then
	assert.Equal(t, []day9.DiskEntry{
		{Id: 0, Start: 0, Length: 4},
		{Id: -1, Start: 4, Length: 5},
		{Id: 1, Start: 9, Length: 1},
	}, dm.Entries)
}

func TestMovesFurthestFileIntoAvailableSpace(t *testing.T) {
	// given
	dm := day9.NewDiskMap(Sample)
	outputs := parsing.Lines(`
0099.111...2...333.44.5555.6666.777.8888..
0099.1117772...333.44.5555.6666.....8888..
0099.111777244.333....5555.6666.....8888..
00992111777.44.333....5555.6666.....8888..`)

	var result bool
	for _, output := range outputs {
		// when
		result = dm.DefragLastWholeFile()

		// then
		assert.True(t, result)
		assert.Equal(t, output, dm.String())
	}

	result = dm.DefragLastWholeFile()
	assert.False(t, result)

	assert.Equal(t, 2858, dm.Checksum())
}

func TestCalculatesChecksumForPart2(t *testing.T) {
	// given
	dm := day9.NewDiskMap(parsing.TrimFile("input.txt"))
	dm.DefragWholeDiskWithWholeFiles()

	// then
	assert.Equal(t, 6307275788409, dm.Checksum())
}
