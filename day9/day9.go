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
	NextFileId   int
	NextLocation int
	Entries      []DiskEntry
}

func NewDiskMap(line string) (dm DiskMap) {
	dm = DiskMap{}

	for _, count := range []rune(line) {
		length, _ := strconv.Atoi(string(count))
		dm.AddEntryFromLength(length)
	}
	return
}

func (dm *DiskMap) AddEntryFromLength(length int) {
	entry := DiskEntry{
		Id:     dm.NextId(),
		Start:  dm.NextLocation,
		Length: length,
	}
	dm.Entries = append(dm.Entries, entry)
	dm.NextLocation += length
}

// NextId is either the next file id for a file entry, or -1 for a space
//
// We increment the next file id as a side effect if we are going to use this one
func (dm *DiskMap) NextId() (id int) {
	nextEntryIsFile := len(dm.Entries)%2 == 0

	if nextEntryIsFile {
		id = dm.NextFileId
		dm.NextFileId++
	} else {
		id = -1
	}
	return
}

// DefragLastFile moves the rightmost file into the first space and returns `true` if anything changed
func (dm *DiskMap) DefragLastFile() bool {
	return dm.DefragFile(dm.LastFile(), dm.FirstSpace())
}

// DefragLastWholeFile moves the rightmost file that can move wholesale into the first space that can hold it
// otherwise it returns false
func (dm *DiskMap) DefragLastWholeFile() bool {
	for filePos, entry := range slices.Backward(dm.Entries) {
		if entry.IsFile() {
			// find first space big enough
			if spacePos := slices.IndexFunc(dm.Entries, func(space DiskEntry) bool {
				return space.IsSpace() && space.Length >= entry.Length && space.Start < entry.Start
			}); spacePos != -1 {
				// found one!
				return dm.DefragFile(filePos, spacePos)
			}
		}
	}

	// none found
	return false
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
		file.Id = -1       // merge the spaces

	// file is longer, leave some behind
	case file.Length > space.Length:
		space.Id = file.Id                         // 'move' part of the file
		file.Length -= space.Length                // adjust the remaining length
		dm.Entries = append(dm.Entries, DiskEntry{ // pad the end
			Id:     -1,
			Start:  file.Start + space.Length,
			Length: space.Length})

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
		file.Id = -1
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

func (dm *DiskMap) FirstSpace() int {
	found := slices.IndexFunc(dm.Entries, func(e DiskEntry) bool { return e.IsSpace() && e.Length > 0 })
	if found == -1 {
		return -1
	}
	return found
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

func (dm *DiskMap) Compact() {
	compacted := make([]DiskEntry, 0)

	for i := 0; i < len(dm.Entries); i++ {
		e := dm.Entries[i]
		if e.IsSpace() {
			for j := i + 1; j < len(dm.Entries) && e.Id == dm.Entries[j].Id; j++ {
				e.Length += dm.Entries[j].Length
				i++
			}
		}
		compacted = append(compacted, e)
	}
	dm.Entries = compacted
}

func (dm *DiskMap) DefragWholeDiskWithWholeFiles() *DiskMap {
	for dm.DefragLastWholeFile() {
		dm.Compact()
	}
	return dm
}
