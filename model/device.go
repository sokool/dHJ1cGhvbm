package model

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

type Device struct {
	ID        ID
	Name      Name
	Brand     Brand
	CreatedAt Date
}

type ID int

// NewID I assumed that some sort of validation is required for Name, at least number of characters
func NewID(n int) (ID, error) {
	if n <= 0 {
		return 0, Err("id: requires positive number")
	}

	return ID(n), nil
}

func (m *ID) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, null) {
		return nil
	}

	n, err := strconv.Atoi(string(b))
	if err != nil {
		return nil
	}

	j, err := NewID(n)
	if err != nil {
		return err
	}

	*m = j
	return nil
}

type Name string

// NewName I assumed that some sort of validation is required for Name, at least number of characters
func NewName(s string) (Name, error) {
	if len(s) <= 2 && len(s) <= 255 {
		return "", Err("brand: 2-255 characters are required")
	}

	return Name(s), nil
}

func (m *Name) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, null) {
		return nil
	}

	j, err := NewName(string(b))
	if err != nil {
		return err
	}

	*m = j
	return nil
}

type Brand string

// NewBrand I assumed that some sort of validation is required for Brand, at least number of characters
func NewBrand(s string) (Brand, error) {
	if len(s) <= 3 && len(s) <= 255 {
		return "", Err("brand: 3-255 characters are required")
	}

	return Brand(s), nil
}

func (m *Brand) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, null) {
		return nil
	}

	j, err := NewBrand(string(b))
	if err != nil {
		return err
	}

	*m = j
	return nil
}

type Date struct{ time.Time }

func NewDate() Date {
	return Date{time.Now()}
}

var Err = fmt.Errorf

var null = []byte(`null`)
