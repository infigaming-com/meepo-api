// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: system/service/v1/system.proto

package v1

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

// Validate checks the field values on Currency with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Currency) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Currency with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CurrencyMultiError, or nil
// if none found.
func (m *Currency) ValidateAll() error {
	return m.validate(true)
}

func (m *Currency) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Currency

	// no validation rules for Enabled

	// no validation rules for Hidden

	// no validation rules for Type

	// no validation rules for Symbol

	// no validation rules for Icon

	// no validation rules for DecimalPlaces

	if len(errors) > 0 {
		return CurrencyMultiError(errors)
	}

	return nil
}

// CurrencyMultiError is an error wrapping multiple validation errors returned
// by Currency.ValidateAll() if the designated constraints aren't met.
type CurrencyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrencyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrencyMultiError) AllErrors() []error { return m }

// CurrencyValidationError is the validation error returned by
// Currency.Validate if the designated constraints aren't met.
type CurrencyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrencyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrencyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrencyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrencyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrencyValidationError) ErrorName() string { return "CurrencyValidationError" }

// Error satisfies the builtin error interface
func (e CurrencyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrency.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrencyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrencyValidationError{}

// Validate checks the field values on AddCurrencyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddCurrencyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddCurrencyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddCurrencyRequestMultiError, or nil if none found.
func (m *AddCurrencyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddCurrencyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCurrency()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddCurrencyRequestValidationError{
					field:  "Currency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddCurrencyRequestValidationError{
					field:  "Currency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCurrency()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddCurrencyRequestValidationError{
				field:  "Currency",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddCurrencyRequestMultiError(errors)
	}

	return nil
}

// AddCurrencyRequestMultiError is an error wrapping multiple validation errors
// returned by AddCurrencyRequest.ValidateAll() if the designated constraints
// aren't met.
type AddCurrencyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddCurrencyRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddCurrencyRequestMultiError) AllErrors() []error { return m }

// AddCurrencyRequestValidationError is the validation error returned by
// AddCurrencyRequest.Validate if the designated constraints aren't met.
type AddCurrencyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddCurrencyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddCurrencyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddCurrencyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddCurrencyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddCurrencyRequestValidationError) ErrorName() string {
	return "AddCurrencyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddCurrencyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddCurrencyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddCurrencyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddCurrencyRequestValidationError{}

// Validate checks the field values on AddCurrencyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddCurrencyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddCurrencyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddCurrencyResponseMultiError, or nil if none found.
func (m *AddCurrencyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddCurrencyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AddCurrencyResponseMultiError(errors)
	}

	return nil
}

// AddCurrencyResponseMultiError is an error wrapping multiple validation
// errors returned by AddCurrencyResponse.ValidateAll() if the designated
// constraints aren't met.
type AddCurrencyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddCurrencyResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddCurrencyResponseMultiError) AllErrors() []error { return m }

// AddCurrencyResponseValidationError is the validation error returned by
// AddCurrencyResponse.Validate if the designated constraints aren't met.
type AddCurrencyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddCurrencyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddCurrencyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddCurrencyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddCurrencyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddCurrencyResponseValidationError) ErrorName() string {
	return "AddCurrencyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddCurrencyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddCurrencyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddCurrencyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddCurrencyResponseValidationError{}

// Validate checks the field values on UpdateCurrencyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCurrencyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCurrencyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCurrencyRequestMultiError, or nil if none found.
func (m *UpdateCurrencyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCurrencyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Currency

	if m.Enabled != nil {
		// no validation rules for Enabled
	}

	if m.Hidden != nil {
		// no validation rules for Hidden
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.Symbol != nil {
		// no validation rules for Symbol
	}

	if m.Icon != nil {
		// no validation rules for Icon
	}

	if m.DecimalPlaces != nil {
		// no validation rules for DecimalPlaces
	}

	if len(errors) > 0 {
		return UpdateCurrencyRequestMultiError(errors)
	}

	return nil
}

// UpdateCurrencyRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateCurrencyRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateCurrencyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCurrencyRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCurrencyRequestMultiError) AllErrors() []error { return m }

// UpdateCurrencyRequestValidationError is the validation error returned by
// UpdateCurrencyRequest.Validate if the designated constraints aren't met.
type UpdateCurrencyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCurrencyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCurrencyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCurrencyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCurrencyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCurrencyRequestValidationError) ErrorName() string {
	return "UpdateCurrencyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCurrencyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCurrencyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCurrencyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCurrencyRequestValidationError{}

// Validate checks the field values on UpdateCurrencyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCurrencyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCurrencyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCurrencyResponseMultiError, or nil if none found.
func (m *UpdateCurrencyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCurrencyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCurrency()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateCurrencyResponseValidationError{
					field:  "Currency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateCurrencyResponseValidationError{
					field:  "Currency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCurrency()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCurrencyResponseValidationError{
				field:  "Currency",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateCurrencyResponseMultiError(errors)
	}

	return nil
}

// UpdateCurrencyResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateCurrencyResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateCurrencyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCurrencyResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCurrencyResponseMultiError) AllErrors() []error { return m }

// UpdateCurrencyResponseValidationError is the validation error returned by
// UpdateCurrencyResponse.Validate if the designated constraints aren't met.
type UpdateCurrencyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCurrencyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCurrencyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCurrencyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCurrencyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCurrencyResponseValidationError) ErrorName() string {
	return "UpdateCurrencyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCurrencyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCurrencyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCurrencyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCurrencyResponseValidationError{}

// Validate checks the field values on GetCurrenciesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCurrenciesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCurrenciesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCurrenciesRequestMultiError, or nil if none found.
func (m *GetCurrenciesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCurrenciesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Enabled != nil {
		// no validation rules for Enabled
	}

	if m.Hidden != nil {
		// no validation rules for Hidden
	}

	if len(errors) > 0 {
		return GetCurrenciesRequestMultiError(errors)
	}

	return nil
}

// GetCurrenciesRequestMultiError is an error wrapping multiple validation
// errors returned by GetCurrenciesRequest.ValidateAll() if the designated
// constraints aren't met.
type GetCurrenciesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCurrenciesRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCurrenciesRequestMultiError) AllErrors() []error { return m }

// GetCurrenciesRequestValidationError is the validation error returned by
// GetCurrenciesRequest.Validate if the designated constraints aren't met.
type GetCurrenciesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCurrenciesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCurrenciesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCurrenciesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCurrenciesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCurrenciesRequestValidationError) ErrorName() string {
	return "GetCurrenciesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCurrenciesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCurrenciesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCurrenciesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCurrenciesRequestValidationError{}

// Validate checks the field values on GetCurrenciesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCurrenciesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCurrenciesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCurrenciesResponseMultiError, or nil if none found.
func (m *GetCurrenciesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCurrenciesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCurrencies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetCurrenciesResponseValidationError{
						field:  fmt.Sprintf("Currencies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetCurrenciesResponseValidationError{
						field:  fmt.Sprintf("Currencies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetCurrenciesResponseValidationError{
					field:  fmt.Sprintf("Currencies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetCurrenciesResponseMultiError(errors)
	}

	return nil
}

// GetCurrenciesResponseMultiError is an error wrapping multiple validation
// errors returned by GetCurrenciesResponse.ValidateAll() if the designated
// constraints aren't met.
type GetCurrenciesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCurrenciesResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCurrenciesResponseMultiError) AllErrors() []error { return m }

// GetCurrenciesResponseValidationError is the validation error returned by
// GetCurrenciesResponse.Validate if the designated constraints aren't met.
type GetCurrenciesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCurrenciesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCurrenciesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCurrenciesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCurrenciesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCurrenciesResponseValidationError) ErrorName() string {
	return "GetCurrenciesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCurrenciesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCurrenciesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCurrenciesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCurrenciesResponseValidationError{}
