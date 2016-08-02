package jutil

import (
	"bytes"
	"io"
)

// QuoteString takes a string as argument and returns the version of that
// string quoted according to the JSON formatting rules.
func QuoteString(s string) string {
	return string(Quote([]byte(s)))
}

// Quote takes a byte slice as argument and returns the copy of that slice
// quoted according to the JSON formatting rules.
func Quote(b []byte) []byte {
	w := &bytes.Buffer{}
	w.Grow(len(b) + 12)
	WriteQuoted(w, b)
	return w.Bytes()
}

// WriteQuoted outputs a byte slice into an io.Writer, quoted according to the
// JSON formatting rules.
func WriteQuoted(w io.Writer, b []byte) (n int, err error) {
	var k int

	k, err = w.Write(quote[:])
	n += k
	if err != nil {
		return
	}

	k, err = WriteEscaped(w, b)
	n += k
	if err != nil {
		return
	}

	k, err = w.Write(quote[:])
	n += k
	return
}

var (
	quote = [...]byte{'"'}
)
