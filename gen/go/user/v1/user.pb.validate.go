// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/v1/user.proto

package user_v1

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

// define the regex for a UUID once up-front
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = UserValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for FirstName

	// no validation rules for LastName

	// no validation rules for Role

	// no validation rules for PhoneNumber

	// no validation rules for Email

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

func (m *User) _validateUuid(uuid string) error {
	if matched := _user_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserRequestMultiError,
// or nil if none found.
func (m *GetUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetUserRequestMultiError(errors)
	}

	return nil
}

// GetUserRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserRequestMultiError) AllErrors() []error { return m }

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserResponseMultiError, or nil if none found.
func (m *GetUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserResponseMultiError(errors)
	}

	return nil
}

// GetUserResponseMultiError is an error wrapping multiple validation errors
// returned by GetUserResponse.ValidateAll() if the designated constraints
// aren't met.
type GetUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserResponseMultiError) AllErrors() []error { return m }

// GetUserResponseValidationError is the validation error returned by
// GetUserResponse.Validate if the designated constraints aren't met.
type GetUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserResponseValidationError) ErrorName() string { return "GetUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserResponseValidationError{}

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FirstName

	// no validation rules for LastName

	// no validation rules for Role

	// no validation rules for PhoneNumber

	// no validation rules for Email

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserResponseMultiError, or nil if none found.
func (m *CreateUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateUserResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateUserResponseMultiError(errors)
	}

	return nil
}

// CreateUserResponseMultiError is an error wrapping multiple validation errors
// returned by CreateUserResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserResponseMultiError) AllErrors() []error { return m }

// CreateUserResponseValidationError is the validation error returned by
// CreateUserResponse.Validate if the designated constraints aren't met.
type CreateUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserResponseValidationError) ErrorName() string {
	return "CreateUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserResponseValidationError{}
