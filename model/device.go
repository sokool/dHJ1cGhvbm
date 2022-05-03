package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Device struct {
	ID        ID
	Name      Name
	Brand     Brand
	CreatedAt Date

	json struct{ Name, Brand string }
}

func (m *Device) UnmarshalJSON(b []byte) error {
	var err error
	if bytes.Equal(b, null) {
		return nil
	}

	if err = json.Unmarshal(b, &m.json); err != nil {
		return err
	}

	if m.Brand, err = NewBrand(m.json.Brand); err != nil {
		return err
	}

	if m.Name, err = NewName(m.json.Name); err != nil {
		return err
	}

	m.CreatedAt = NewDate()

	return nil
}

type ID int

// NewID I assumed that some sort of validation is required for Name, at least number of characters
func NewID(n int) (ID, error) {
	if n <= 0 {
		return 0, Err("id: requires positive number")
	}

	return ID(n), nil
}

func NewParseID(s string) (ID, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return NewID(n)
}

func (m ID) IsZero() bool { return m <= 0 }

func (m *ID) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, null) {
		return nil
	}

	n, err := strconv.Atoi(string(b))
	if err != nil {
		return err
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
	if s = strings.TrimSpace(s); len(s) < 2 || len(s) >= 255 {
		return "", Err("name: 2-255 characters are required")
	}

	return Name(s), nil
}

func (m Name) IsZero() bool { return m == "" }

func (m *Name) UnmarshalJSON(b []byte) error {
	s := len(b)
	if bytes.Equal(b, null) || s <= 2 {
		return nil
	}

	j, err := NewName(string(b[1 : s-1]))
	if err != nil {
		return err
	}

	*m = j
	return nil
}

type Brand string

// NewBrand I assumed that some sort of validation is required for Brand, at least number of characters
func NewBrand(s string) (Brand, error) {
	if s = strings.TrimSpace(s); len(s) < 3 || len(s) >= 255 {
		return "", Err("brand: 3-255 characters are required")
	}

	return Brand(s), nil
}

func (m Brand) IsZero() bool { return m == "" }

func (m *Brand) UnmarshalJSON(b []byte) error {
	s := len(b)
	if bytes.Equal(b, null) || s <= 2 {
		return nil
	}

	j, err := NewBrand(string(b[1 : s-1]))
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
