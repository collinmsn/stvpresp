package resp

import (
	"bytes"
)

// parseLenLine takes a valid RESP array or bulk string length specification
// and returns the specified length. It returns -1 for null bulk strings. If
// the line is invalid, an error will be returned.
func parseLenLine(line []byte) (int, error) {
	if len(line) < 4 {
		return -1, ErrSyntaxError
	}

	// Shortcut for null bulk strings
	if len(line) == 5 && line[1] == '-' && line[2] == '1' {
		return -1, nil
	}

	var n int
	for _, b := range line[1 : len(line)-2] {
		n *= 10
		if b < '0' || b > '9' {
			return -1, ErrSyntaxError
		}
		n += int(b - '0')
	}

	return n, nil
}

// indexLineEnd returns the index of the final character of the first line in
// the given RESP byte slice. If no valid line ending can be found, it returns
// -1.
func indexLineEnd(slice []byte) int {
	i := bytes.IndexByte(slice, '\n')
	if i > 0 && slice[i-1] == '\r' {
		return i
	}
	return -1
}
