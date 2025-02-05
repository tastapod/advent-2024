package day9

import (
	"slices"
	"strconv"
	"strings"
)

type DiskEntry struct {
	Id            int // -1 for spaces
	Start, Length int
}

func (e *DiskEntry) IsSpace() bool {
	return e.Id == -1
}

func (e *DiskEntry) IsFile() bool {
	return !e.IsSpace()
}

func (e *DiskEntry) Checksum() (result int) {
	if e.IsSpace() {
		return 0
	}
	for i := e.Start; i < e.Start+e.Length; i++ {
		result += i * e.Id
	}
	return
}

type DiskMap struct {
	Entries []DiskEntry
}

func NewDiskMap(line string) DiskMap {
	start := 0
	nextId := 0
	entries := make([]DiskEntry, 0)

	for i, length := range []rune(line) {
		length, _ := strconv.Atoi(string(length))
		if length == 0 {
			continue
		}
		isFile := i%2 == 0
		var id int
		if isFile {
			id = nextId
			nextId++
		} else {
			id = -1
		}
		entries = append(entries, DiskEntry{Id: id, Start: start, Length: length})
		start += length
	}
	return DiskMap{Entries: entries}
}

// DefragLastFile moves the rightmost file into the first space and returns `true` if anything changed
func (dm *DiskMap) DefragLastFile() bool {
	filePos := dm.LastFile()
	return dm.DefragFile(filePos, dm.FirstSpace(filePos, 1))
}

func (dm *DiskMap) DefragFile(filePos int, spacePos int) bool {
	// are we done?
	if filePos == -1 || spacePos == -1 {
		// all finished
		return false
	}

	file := &dm.Entries[filePos]
	space := &dm.Entries[spacePos]

	if space.Start > file.Start {
		// first space is after last file
		return false
	}

	switch {
	// same size
	case file.Length == space.Length:
		space.Id = file.Id // 'move' the file
		dm.MergeSpaces(filePos)

	// file is longer, leave some behind
	case file.Length > space.Length:
		space.Id = file.Id                         // 'move' enough file to fill the space
		file.Length -= space.Length                // adjust the remaining length
		dm.Entries = append(dm.Entries, DiskEntry{ // pad the end
			Id:     -1,
			Start:  file.Start + space.Length,
			Length: space.Length})
		dm.MergeSpaces(len(dm.Entries) - 1)

	// file is shorter
	default:
		// insert file before space and shorten the space
		movedFile := DiskEntry{Id: file.Id, Start: space.Start, Length: file.Length}
		dm.Entries = slices.Insert(dm.Entries, spacePos, movedFile)

		// file and space positions have moved
		filePos++
		file = &dm.Entries[filePos]

		spacePos++
		space = &dm.Entries[spacePos]

		// shrink space
		space.Start += movedFile.Length
		space.Length -= movedFile.Length

		// blank out original file
		dm.MergeSpaces(filePos)
	}
	return true
}

func (dm *DiskMap) LastFile() int {
	for i := len(dm.Entries) - 1; i >= 0; i-- {
		entry := &dm.Entries[i]
		if entry.IsFile() {
			return i
		}
	}
	return -1
}

func (dm *DiskMap) FirstSpace(filePos, minLength int) int {
	spacePos := slices.IndexFunc(dm.Entries, func(e DiskEntry) bool { return e.IsSpace() && e.Length >= minLength })
	if spacePos > filePos {
		return -1
	}
	return spacePos
}

func (dm *DiskMap) DefragWholeDisk() *DiskMap {
	for dm.DefragLastFile() {
	}
	return dm
}

func (dm *DiskMap) String() string {
	var sb strings.Builder

	for _, e := range dm.Entries {
		var symbol string
		if e.IsFile() {
			symbol = strconv.Itoa(e.Id)
		} else {
			symbol = "."
		}
		sb.WriteString(strings.Repeat(symbol, e.Length))
	}
	return sb.String()
}

func (dm *DiskMap) Checksum() (result int) {
	for _, e := range dm.Entries {
		result += e.Checksum()
	}
	return
}

func (dm *DiskMap) DefragWholeDiskWithWholeFiles() *DiskMap {
	// step backwards through all files
	maxId := dm.MaxId()
	for id := maxId; id >= 0; id-- {
		//debug.Debug(dm.Entries)
		filePos := slices.IndexFunc(dm.Entries, func(e DiskEntry) bool { return e.Id == id })
		if spacePos := dm.FirstSpace(filePos, dm.Entries[filePos].Length); spacePos != -1 {
			dm.DefragFile(filePos, spacePos)
			//debug.Debug(dm.String())
		}
	}
	return dm
}

// MergeSpaces checks for spaces either side of current position and merges where possible
func (dm *DiskMap) MergeSpaces(pos int) {

	current := &dm.Entries[pos]
	var previous, next *DiskEntry

	if pos > 0 {
		previous = &dm.Entries[pos-1]
	}
	if pos < len(dm.Entries)-1 {
		next = &dm.Entries[pos+1]
	}

	if previous != nil && previous.IsSpace() {
		// merge with previous
		if next != nil && next.IsSpace() {
			// also merge with next
			dm.Entries = slices.Replace(dm.Entries, pos-1, pos+2, DiskEntry{
				Id:     -1,
				Start:  previous.Start,
				Length: previous.Length + current.Length + next.Length,
			})
		} else {
			// just previous and current
			dm.Entries = slices.Replace(dm.Entries, pos-1, pos+1, DiskEntry{
				Id:     -1,
				Start:  previous.Start,
				Length: previous.Length + current.Length,
			})
		}
	} else if next != nil && next.IsSpace() {
		// just current and next
		dm.Entries = slices.Replace(dm.Entries, pos, pos+2, DiskEntry{
			Id:     -1,
			Start:  current.Start,
			Length: current.Length + next.Length,
		})
	} else {
		// nope, just us
		current.Id = -1
	}
}

func (dm *DiskMap) MaxId() int {
	for _, entry := range slices.Backward(dm.Entries) {
		if entry.IsFile() {
			return entry.Id
		}
	}
	return -1
}
