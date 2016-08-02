package jutil

import (
	"reflect"
	"strings"
)

// Tag represents the result of parsing the json tag of a struct field.
type Tag struct {

	// Name is the field name that should be used when serializing JSON.
	Name string

	// Omitempty is true if the struct field json tag had `omitempty` set.
	Omitempty bool

	// Skip is true if the struct field json tag started with `-`.
	Skip bool
}

// ParseStructField parses the tag of a struct field that may or may not
// have a `json` tag set, returing the result as a Tag field.
func ParseStructField(f reflect.StructField) Tag {
	t := ParseStructTag(f.Tag.Get("json"))
	if len(t.Name) == 0 {
		t.Name = f.Name
	}
	return t
}

// ParseStructTag parses a raw json tag obtained from a struct field,
// returining the results as a Tag value.
func ParseStructTag(tag string) Tag {
	name, tag := parseNextTagToken(tag)
	token, _ := parseNextTagToken(tag)
	return Tag{
		Name:      name,
		Skip:      name == "-",
		Omitempty: token == "omitempty",
	}
}

func parseNextTagToken(tag string) (token string, next string) {
	if split := strings.IndexByte(tag, ','); split < 0 {
		token = tag
	} else {
		token, next = tag[:split], tag[split+1:]
	}
	return
}
