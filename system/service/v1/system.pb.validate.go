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

// Validate checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckRequestMultiError, or nil if none found.
func (m *HealthCheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return HealthCheckRequestMultiError(errors)
	}

	return nil
}

// HealthCheckRequestMultiError is an error wrapping multiple validation errors
// returned by HealthCheckRequest.ValidateAll() if the designated constraints
// aren't met.
type HealthCheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckRequestMultiError) AllErrors() []error { return m }

// HealthCheckRequestValidationError is the validation error returned by
// HealthCheckRequest.Validate if the designated constraints aren't met.
type HealthCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckRequestValidationError) ErrorName() string {
	return "HealthCheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckRequestValidationError{}

// Validate checks the field values on HealthCheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckResponseMultiError, or nil if none found.
func (m *HealthCheckResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return HealthCheckResponseMultiError(errors)
	}

	return nil
}

// HealthCheckResponseMultiError is an error wrapping multiple validation
// errors returned by HealthCheckResponse.ValidateAll() if the designated
// constraints aren't met.
type HealthCheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckResponseMultiError) AllErrors() []error { return m }

// HealthCheckResponseValidationError is the validation error returned by
// HealthCheckResponse.Validate if the designated constraints aren't met.
type HealthCheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckResponseValidationError) ErrorName() string {
	return "HealthCheckResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckResponseValidationError{}

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

	// no validation rules for Currency

	// no validation rules for DecimalPlaces

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

	if all {
		switch v := interface{}(m.GetCurrency()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddCurrencyResponseValidationError{
					field:  "Currency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddCurrencyResponseValidationError{
					field:  "Currency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCurrency()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddCurrencyResponseValidationError{
				field:  "Currency",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

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

// Validate checks the field values on CurrenciesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CurrenciesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CurrenciesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CurrenciesRequestMultiError, or nil if none found.
func (m *CurrenciesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CurrenciesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Currency

	if len(errors) > 0 {
		return CurrenciesRequestMultiError(errors)
	}

	return nil
}

// CurrenciesRequestMultiError is an error wrapping multiple validation errors
// returned by CurrenciesRequest.ValidateAll() if the designated constraints
// aren't met.
type CurrenciesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrenciesRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrenciesRequestMultiError) AllErrors() []error { return m }

// CurrenciesRequestValidationError is the validation error returned by
// CurrenciesRequest.Validate if the designated constraints aren't met.
type CurrenciesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrenciesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrenciesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrenciesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrenciesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrenciesRequestValidationError) ErrorName() string {
	return "CurrenciesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CurrenciesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrenciesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrenciesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrenciesRequestValidationError{}

// Validate checks the field values on CurrenciesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CurrenciesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CurrenciesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CurrenciesResponseMultiError, or nil if none found.
func (m *CurrenciesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CurrenciesResponse) validate(all bool) error {
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
					errors = append(errors, CurrenciesResponseValidationError{
						field:  fmt.Sprintf("Currencies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CurrenciesResponseValidationError{
						field:  fmt.Sprintf("Currencies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CurrenciesResponseValidationError{
					field:  fmt.Sprintf("Currencies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CurrenciesResponseMultiError(errors)
	}

	return nil
}

// CurrenciesResponseMultiError is an error wrapping multiple validation errors
// returned by CurrenciesResponse.ValidateAll() if the designated constraints
// aren't met.
type CurrenciesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrenciesResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrenciesResponseMultiError) AllErrors() []error { return m }

// CurrenciesResponseValidationError is the validation error returned by
// CurrenciesResponse.Validate if the designated constraints aren't met.
type CurrenciesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrenciesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrenciesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrenciesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrenciesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrenciesResponseValidationError) ErrorName() string {
	return "CurrenciesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CurrenciesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrenciesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrenciesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrenciesResponseValidationError{}
