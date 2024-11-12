// Code generated by mfd-generator v0.4.5; DO NOT EDIT.

//nolint:all
//lint:file-ignore U1000 ignore unused code, it's generated
package db

import (
	"unicode/utf8"
)

const (
	ErrEmptyValue = "empty"
	ErrMaxLength  = "len"
	ErrWrongValue = "value"
)

func (c City) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if utf8.RuneCountInString(c.Name) > 250 {
		errors[Columns.City.Name] = ErrMaxLength
	}

	return errors, len(errors) == 0
}

func (d Direction) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if utf8.RuneCountInString(d.Code) > 255 {
		errors[Columns.Direction.Code] = ErrMaxLength
	}

	if utf8.RuneCountInString(d.Name) > 250 {
		errors[Columns.Direction.Name] = ErrMaxLength
	}

	return errors, len(errors) == 0
}

func (u Universyty) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if utf8.RuneCountInString(u.Name) > 250 {
		errors[Columns.Universyty.Name] = ErrMaxLength
	}

	return errors, len(errors) == 0
}

func (u User) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	return errors, len(errors) == 0
}