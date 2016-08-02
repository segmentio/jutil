package jutil

import "testing"

func TestQuoteString(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "",
			out: `""`,
		},
		{
			in:  "Hello World!",
			out: `"Hello World!"`,
		},
		{
			in:  "Hello\"World!",
			out: `"Hello\"World!"`,
		},
		{
			in:  "Hello/World!",
			out: `"Hello\/World!"`,
		},
		{
			in:  "Hello\\World!",
			out: `"Hello\\World!"`,
		},
		{
			in:  "Hello\nWorld!",
			out: `"Hello\nWorld!"`,
		},
		{
			in:  "Hello\tWorld!",
			out: `"Hello\tWorld!"`,
		},
		{
			in:  "Hello\rWorld!",
			out: `"Hello\rWorld!"`,
		},
		{
			in:  "Hello\vWorld!",
			out: `"Hello\vWorld!"`,
		},
		{
			in:  "Hello\bWorld!",
			out: `"Hello\bWorld!"`,
		},
		{
			in:  "Hello\fWorld!",
			out: `"Hello\fWorld!"`,
		},
	}

	for _, test := range tests {
		if s := QuoteString(test.in); s != test.out {
			t.Errorf("%#v: invalid quoted string: %#v != %#v", test.in, test.out, s)
		}
	}
}
