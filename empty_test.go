package jutil

import "testing"

func TestIsEmptyTrue(t *testing.T) {
	tests := []interface{}{
		(*int)(nil),
		false,
		0,
		0.0,
		"",
		[]byte{},
		[]int{},
		map[string]interface{}{},
	}

	for _, test := range tests {
		if !IsEmptyValue(test) {
			t.Errorf("%#v should be an empty value", test)
		}
	}
}

func TestIsEmptyFale(t *testing.T) {
	tests := []interface{}{
		true,
		1,
		1.0,
		"Hello World!",
		[]byte{0},
		[]int{0},
		map[string]interface{}{"A": 1},
		struct{}{},
		struct{ A int }{},
		struct{ A int }{42},
	}

	for _, test := range tests {
		if IsEmptyValue(test) {
			t.Errorf("%#v should not be an empty value", test)
		}
	}
}
