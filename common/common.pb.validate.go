// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/common.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Id with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Id) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Id with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IdMultiError, or nil if none found.
func (m *Id) ValidateAll() error {
	return m.validate(true)
}

func (m *Id) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IdMultiError(errors)
	}

	return nil
}

// IdMultiError is an error wrapping multiple validation errors returned by
// Id.ValidateAll() if the designated constraints aren't met.
type IdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdMultiError) AllErrors() []error { return m }

// IdValidationError is the validation error returned by Id.Validate if the
// designated constraints aren't met.
type IdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdValidationError) ErrorName() string { return "IdValidationError" }

// Error satisfies the builtin error interface
func (e IdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdValidationError{}

// Validate checks the field values on Meta with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Meta) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Meta with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MetaMultiError, or nil if none found.
func (m *Meta) ValidateAll() error {
	return m.validate(true)
}

func (m *Meta) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetaValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "CreatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "CreatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetaValidationError{
				field:  "CreatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetaValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "UpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "UpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetaValidationError{
				field:  "UpdatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDeletedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetaValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetaValidationError{
				field:  "DeletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MetaMultiError(errors)
	}

	return nil
}

// MetaMultiError is an error wrapping multiple validation errors returned by
// Meta.ValidateAll() if the designated constraints aren't met.
type MetaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetaMultiError) AllErrors() []error { return m }

// MetaValidationError is the validation error returned by Meta.Validate if the
// designated constraints aren't met.
type MetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetaValidationError) ErrorName() string { return "MetaValidationError" }

// Error satisfies the builtin error interface
func (e MetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetaValidationError{}
