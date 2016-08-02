package jutil

import (
	"reflect"
	"testing"
)

func TestParseJsonStructTag(t *testing.T) {
	tests := []struct {
		tag string
		res Tag
	}{
		{
			tag: "",
			res: Tag{},
		},
		{
			tag: "hello",
			res: Tag{Name: "hello"},
		},
		{
			tag: ",omitempty",
			res: Tag{Omitempty: true},
		},
		{
			tag: "-",
			res: Tag{Name: "-", Skip: true},
		},
		{
			tag: "hello,omitempty",
			res: Tag{Name: "hello", Omitempty: true},
		},
		{
			tag: "-,omitempty",
			res: Tag{Name: "-", Omitempty: true, Skip: true},
		},
	}

	for _, test := range tests {
		if res := ParseJsonStructTag(test.tag); res != test.res {
			t.Errorf("%s: %#v != %#v", test.tag, test.res, res)
		}
	}
}

func TestParseJsonStructField(t *testing.T) {
	tests := []struct {
		val interface{}
		res Tag
	}{
		{
			val: struct{ F int }{},
			res: Tag{Name: "F"},
		},
		{
			val: struct {
				F int `json:"f"`
			}{},
			res: Tag{Name: "f"},
		},
		{
			val: struct {
				F int `json:"-"`
			}{},
			res: Tag{Name: "-", Skip: true},
		},
		{
			val: struct {
				F int `json:",omitempty"`
			}{},
			res: Tag{Name: "F", Omitempty: true},
		},
	}

	for _, test := range tests {
		if res := ParseJsonStructField(reflect.TypeOf(test.val).Field(0)); res != test.res {
			t.Errorf("%s: %#v != %#v", test.val, test.res, res)
		}
	}
}
