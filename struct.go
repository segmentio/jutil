package jutil

import (
	"reflect"
	"sync"
)

// Struct is used to represent a Go structure in internal data structures that
// cache meta information to make field lookups faster and avoid having to use
// reflection to lookup the same type information over and over again.
type Struct []StructField

// LookupStruct behaves like MakeStruct but uses a global cache to avoid having
// to recreate the struct values when not needed.
//
// As much as possible you should be using this function instead of calling
// MakeStruct or maintaining your own cache so the program can efficiently make
// use of the cache and avoid storing duplicate information in different parts
// of the program.
func LookupStruct(t reflect.Type) Struct {
	return structCache.Lookup(t)
}

// MakeStruct takes a Go type as argument and extract information to make a new
// Struct value.
// The type has to be a struct type or a panic will be raised.
func MakeStruct(t reflect.Type) Struct {
	n := t.NumField()
	s := make(Struct, 0, n)

	for i := 0; i != n; i++ {
		if f := MakeStructField(t.Field(i)); !f.Skip {
			s = append(s, f)
		}
	}

	return s
}

// StructField represents a single field of a struct and carries information
// useful to the algorithms of the jutil package.
type StructField struct {
	// The index of the field in the structure.
	Index []int

	// The name of the field once serialized to JSON.
	Name string

	// True if the field has to be omitted when it has an empty value.
	Omitempty bool

	// True if the field should be skipped entirely.
	Skip bool
}

// MakeStructField takes a Go struct field as argument argument and returns its
// StructType representation.
func MakeStructField(f reflect.StructField) StructField {
	tag := ParseStructField(f)

	field := StructField{
		Index:     f.Index,
		Name:      tag.Name,
		Omitempty: tag.Omitempty,
		Skip:      tag.Skip,
	}

	if len(f.PkgPath) != 0 && !f.Anonymous { // unexported
		field.Skip = true
	}

	return field
}

// StructCache is a simple cache for mapping Go types to Struct values.
type StructCache struct {
	mutex sync.RWMutex
	store map[reflect.Type]Struct
}

// NewStructCache creates and returns a new StructCache value.
func NewStructCache() *StructCache {
	return &StructCache{
		store: make(map[reflect.Type]Struct),
	}
}

// Lookup takes a Go type as argument and returns the matching Struct value,
// potentially creating it if it didn't already exist.
// This method is safe to call from multiple goroutines.
func (cache *StructCache) Lookup(t reflect.Type) (s Struct) {
	cache.mutex.RLock()
	s = cache.store[t]
	cache.mutex.RUnlock()

	if s == nil {
		s = MakeStruct(t)
		cache.mutex.Lock()
		cache.store[t] = s
		cache.mutex.Unlock()
	}

	return
}

var (
	// This struct cache is used to avoid reusing reflection over and over when
	// the jutil functions are called. The performance improvements on iterating
	// over struct fields are huge, this is a really important optimization:
	//
	// benchmark                                   old ns/op     new ns/op     delta
	// BenchmarkLengthStructZero                   53.9          99.9          +85.34%
	// BenchmarkLengthStructNonZero                746           411           -44.91%
	// BenchmarkLengthStructOmitEmptyZero          779           174           -77.66%
	// BenchmarkLengthStructOmpitemptytNonZero     1119          425           -62.02%
	//
	// Note: Disregard the performance loss on the `StructZero` benchmark, this
	// is testing an empty struct with no field, which is just a baseline and not
	// actually useful in real-world use cases.
	//
	structCache = NewStructCache()
)
