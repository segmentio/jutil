package jutil

import (
	"bytes"
	"io"
)

// EscapeString takes a string as argument and returns the version of that
// string where every special characters have been escaped according to the
// JSON formatting rules.
func EscapeString(s string) string {
	return string(Escape([]byte(s)))
}

// Escape takes a byte slice as argument and returns the copy of that slice
// where every special characters have been escaped according to the JSON
// formatting rules.
func Escape(b []byte) []byte {
	w := &bytes.Buffer{}
	w.Grow(len(b) + 10)
	WriteEscaped(w, b)
	return w.Bytes()
}

// WriteEscaped outputs a byte slice into an io.Writer where every special
// character has been escaped according to the JSON formatting rules.
func WriteEscaped(w io.Writer, b []byte) (n int, err error) {
	var e [1]byte
	var k int

	if len(b) < 100 {
		// Fast path optimization for short strings that don't have any
		// characters to escape.
		const escapedBytes = "/\"\\\n\t\r\v\b\f"

		switch len(b) {
		case 0:
		case 1:
			for i := range escapedBytes {
				if b[0] == escapedBytes[i] {
					goto loop
				}
			}
		default:
			for i := range escapedBytes {
				if bytes.IndexByte(b, escapedBytes[i]) >= 0 {
					goto loop
				}
			}
		}

		return w.Write(b)
	}

loop:
	for len(b) != 0 {
		for i, c := range b {

			switch c {
			case '"':
				e[0] = '"'

			case '/':
				e[0] = '/'

			case '\\':
				e[0] = '\\'

			case '\n':
				e[0] = 'n'

			case '\t':
				e[0] = 't'

			case '\r':
				e[0] = 'r'

			case '\v':
				e[0] = 'v'

			case '\b':
				e[0] = 'b'

			case '\f':
				e[0] = 'f'
			}

			if e[0] != 0 {
				k, err = w.Write(b[:i])
				n += k
				if err != nil {
					return
				}

				k, err = w.Write(escape[:])
				n += k
				if err != nil {
					return
				}

				k, err = w.Write(e[:])
				n += k
				if err != nil {
					return
				}

				e[0] = 0
				b = b[i+1:]
				continue loop
			}
		}

		k, err = w.Write(b)
		n += k
		break
	}

	return
}

var (
	escape = [...]byte{'\\'}
)
