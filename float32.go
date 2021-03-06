// Code generated by go generate; DO NOT EDIT.
// This file was generated by scripts/models/gen_nullable_types.go
package nullable

import (
	"bytes"
	"encoding/json"
	"github.com/go-openapi/strfmt"
)

// Float32 represents a float32 that may be null or not
// present in json at all.
type Float32 struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid float32
	Value   float32
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (f *Float32) Ptr() *float32 {
	if f.Present && f.Valid {
		return &f.Value
	}

	return nil
}

// UnmarshalJSON implements json.Marshaler interface.
func (f *Float32) UnmarshalJSON(data []byte) error {
	f.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &f.Value); err != nil {
		return err
	}

	f.Valid = true
	return nil
}

// Validate implements runtime.Validateable interface for go-swagger generation.
func (f *Float32) Validate(formats strfmt.Registry) error {
	return nil
}

// Float32Slice represents a []float32 that may be null or not
// present in json at all.
type Float32Slice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []float32
}

// UnmarshalJSON implements json.Marshaler interface.
func (f *Float32Slice) UnmarshalJSON(data []byte) error {
	f.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &f.Value); err != nil {
		return err
	}

	f.Valid = true
	return nil
}
