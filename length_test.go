package jutil

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/segmentio/ecs-logs-go"
)

const longString = `"Package json implements encoding and decoding of JSON objects as defined in RFC 4627. The mapping between JSON objects and Go values is described in the documentation for the Marshal and Unmarshal functions.")`

func TestLength(t *testing.T) {
	tests := []interface{}{
		nil,

		true,
		false,

		0,
		1,
		42,
		-1,
		-42,
		0.1234,

		"",
		"Hello World!",
		"Hello\nWorld!",

		[]byte(""),
		[]byte("Hello World!"),

		json.Number("0"),
		json.Number("1.2345"),

		time.Now(),
		12 * time.Hour,

		[]int{},
		[]int{1, 2, 3},
		[]string{"hello", "world"},
		[]interface{}{nil, true, 42, "hey!"},

		map[string]string{},
		map[string]int{"answer": 42},
		map[string]interface{}{
			"A": nil,
			"B": true,
			"C": 42,
			"D": "hey!",
		},

		struct{}{},
		struct{ Answer int }{42},
		struct {
			A int
			B int
			C int
		}{1, 2, 3},
		struct {
			Question string
			Answer   string
		}{"How are you?", "Well"},

		map[string]interface{}{
			"struct": struct {
				OK bool `json:",omitempty"`
			}{false},
			"what?": struct {
				List   []interface{}
				String string
			}{
				List:   []interface{}{1, 2, 3},
				String: "Hello World!",
			},
		},

		ecslogs.Event{
			Level:   ecslogs.DEBUG,
			Time:    time.Now(),
			Info:    ecslogs.EventInfo{Host: "localhost"},
			Data:    ecslogs.EventData{"hello": "world"},
			Message: "Hello World!",
		},
	}

	for _, test := range tests {
		b, _ := json.Marshal(test)

		if n, err := Length(test); err != nil {
			t.Errorf("%#v => %s", test, err)
		} else if n != len(b) {
			t.Errorf("%#v => %d != %d (%s)", test, n, len(b), string(b))
		}
	}
}

func benchLength(b *testing.B, v interface{}) {
	for i := 0; i != b.N; i++ {
		benchLengthFunc(v)
	}
}

func BenchmarkLengthBoolZero(b *testing.B) {
	benchLength(b, false)
}

func BenchmarkLengthBoolNonZero(b *testing.B) {
	benchLength(b, true)
}

func BenchmarkLengthIntZero(b *testing.B) {
	benchLength(b, 0)
}

func BenchmarkLengthIntNonZero(b *testing.B) {
	benchLength(b, 1234567890)
}

func BenchmarkLengthFloatZero(b *testing.B) {
	benchLength(b, 0.0)
}

func BenchmarkLengthFloatNonZero(b *testing.B) {
	benchLength(b, 12345.67890)
}

func BenchmarkLengthStringZero(b *testing.B) {
	benchLength(b, "")
}

func BenchmarkLengthStringNonZero(b *testing.B) {
	benchLength(b, longString)
}

func BenchmarkLengthBytesZero(b *testing.B) {
	benchLength(b, []byte{})
}

func BenchmarkLengthBytesNonZero(b *testing.B) {
	benchLength(b, []byte(longString))
}

func BenchmarkLengthSliceInterfaceZero(b *testing.B) {
	benchLength(b, []interface{}{})
}

func BenchmarkLengthSliceInterfaceNonZero(b *testing.B) {
	benchLength(b, []interface{}{})
}

func BenchmarkLengthSliceBoolZero(b *testing.B) {
	benchLength(b, []bool{})
}

func BenchmarkLengthSliceBoolNonZero(b *testing.B) {
	benchLength(b, []bool{
		true, false, true, false, true, false, true, false, true, false, true, false,
	})
}

func BenchmarkLengthTimeZero(b *testing.B) {
	benchLength(b, time.Time{})
}

func BenchmarkLengthTimeNonZero(b *testing.B) {
	benchLength(b, time.Now())
}

func BenchmarkLengthDurationZero(b *testing.B) {
	benchLength(b, time.Duration(0))
}

func BenchmarkLengthDurationNonZero(b *testing.B) {
	benchLength(b, 12*time.Hour)
}

func BenchmarkLengthMapStringInterfaceZero(b *testing.B) {
	benchLength(b, map[string]interface{}{})
}

func BenchmarkLengthMapStringInterfaceNonZero(b *testing.B) {
	benchLength(b, map[string]interface{}{
		"0": true,
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
	})
}

func BenchmarkLengthMapStringStringZero(b *testing.B) {
	benchLength(b, map[string]string{})
}

func BenchmarkLengthMapStringStringNonZero(b *testing.B) {
	benchLength(b, map[string]string{
		"0": "",
		"1": "",
		"2": "",
		"3": "",
		"4": "",
		"5": "",
		"6": "",
		"7": "",
		"8": "",
		"9": "",
	})
}

func BenchmarkLengthStructZero(b *testing.B) {
	benchLength(b, struct{}{})
}

func BenchmarkLengthStructNonZero(b *testing.B) {
	benchLength(b, struct {
		A int
		B int
		C int
	}{1, 2, 3})
}

func BenchmarkLengthStructOmitEmptyZero(b *testing.B) {
	benchLength(b, struct {
		A int `json:",omitempty"`
		B int `json:",omitempty"`
		C int `json:",omitempty"`
	}{})
}

func BenchmarkLengthStructOmpitemptytNonZero(b *testing.B) {
	benchLength(b, struct {
		A int `json:",omitempty"`
		B int `json:",omitempty"`
		C int `json:",omitempty"`
	}{1, 2, 3})
}
