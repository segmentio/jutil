package jutil

import (
	"reflect"
	"testing"
)

func TestMakeStructField(t *testing.T) {
	tests := []struct {
		s reflect.StructField
		f StructField
	}{
		{
			s: reflect.TypeOf(struct{ A int }{}).Field(0),
			f: StructField{
				Index:     []int{0},
				Name:      "A",
				Omitempty: false,
				Skip:      false,
			},
		},
		{
			s: reflect.TypeOf(struct{ a int }{}).Field(0),
			f: StructField{
				Index:     []int{0},
				Name:      "a",
				Omitempty: false,
				Skip:      true,
			},
		},
	}

	for _, test := range tests {
		if f := MakeStructField(test.s); !reflect.DeepEqual(test.f, f) {
			t.Errorf("%#v != %#v", test.f, f)
		}
	}
}
