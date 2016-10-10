package filesystem

const (
	FileModeDefault fileMode = 1 << iota
	FileModeRead
	FileModeWrite
	FileModeAppend
	FileModeBinary
)

const (
	CharsetDefault Charset = ""
	CharsetASCII           = "US-ASCII"
	CharsetUTF8            = "UTF-8"
)

type fileMode int

type Charset string

func parseFileMode(mode fileMode) string {
	var res string
	switch {
	case mode&FileModeRead != 0:
		res = "r"

	case mode&FileModeWrite != 0:
		res = "w"

	case mode&FileModeAppend != 0:
		res = "a"

	}

	if mode&FileModeBinary != 0 {
		res += "b"
	}

	return res
}
