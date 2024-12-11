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

// DefragLastFile moves the hindmost file into the first space and returns `true` if anything changed
func (dm *DiskMap) DefragLastFile() bool {
	_, file := dm.LastFile()
	spacePos, space := dm.FirstSpace()

	// are we done?
	if file == nil || space == nil || space.Start > file.Start {
		// first space is beyond the last file!
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
		// replace space with file+space
		dm.Entries = slices.Replace(dm.Entries, spacePos, spacePos+1,
			DiskEntry{Id: file.Id, Start: space.Start, Length: file.Length},
			DiskEntry{Id: -1, Start: space.Start + file.Length, Length: space.Length - file.Length},
		)
		// blank file at end (NOTE: file no longer points to the correct entry!)
		_, file = dm.LastFile()
		file.Id = -1
	}
	return true
}

func (dm *DiskMap) LastFile() (int, *DiskEntry) {
	for i := len(dm.Entries) - 1; i >= 0; i-- {
		entry := &dm.Entries[i]
		if entry.IsFile() {
			return i, entry
		}
	}
	return -1, nil
}

func (dm *DiskMap) FirstSpace() (int, *DiskEntry) {
	found := slices.IndexFunc(dm.Entries, func(e DiskEntry) bool { return e.IsSpace() && e.Length > 0 })
	if found == -1 {
		return -1, nil
	}
	return found, &dm.Entries[found]
}

func (dm *DiskMap) DefragWholeDisk() {
	for dm.DefragLastFile() {
	}
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
