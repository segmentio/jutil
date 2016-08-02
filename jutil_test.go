package jutil

import (
	"encoding/json"
	"flag"
	"os"
	"testing"
)

var benchLengthFunc func(interface{})

func TestMain(m *testing.M) {
	var benchLength string

	flag.StringVar(&benchLength, "bench-length", "Length", "the length function to benchmark")
	flag.Parse()

	switch benchLength {
	case "json.Marshal":
		benchLengthFunc = func(v interface{}) { json.Marshal(v) }

	default:
		benchLengthFunc = func(v interface{}) { Length(v) }
	}

	os.Exit(m.Run())
}
