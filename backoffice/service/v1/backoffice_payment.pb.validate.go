// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: backoffice/service/v1/backoffice_payment.proto

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

// Validate checks the field values on TransactionInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TransactionInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TransactionInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TransactionInfoMultiError, or nil if none found.
func (m *TransactionInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *TransactionInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TransactionId

	// no validation rules for PaTransactionId

	// no validation rules for GatewayTransactionId

	// no validation rules for OperatorId

	// no validation rules for UserId

	// no validation rules for Vip

	// no validation rules for Amount

	// no validation rules for Currency

	// no validation rules for Fee

	// no validation rules for PaymentMethod

	// no validation rules for PaymentChannel

	// no validation rules for Protocol

	// no validation rules for Type

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TransactionInfoValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TransactionInfoValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TransactionInfoValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TransactionInfoValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TransactionInfoValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TransactionInfoValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TransactionInfoMultiError(errors)
	}

	return nil
}

// TransactionInfoMultiError is an error wrapping multiple validation errors
// returned by TransactionInfo.ValidateAll() if the designated constraints
// aren't met.
type TransactionInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransactionInfoMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransactionInfoMultiError) AllErrors() []error { return m }

// TransactionInfoValidationError is the validation error returned by
// TransactionInfo.Validate if the designated constraints aren't met.
type TransactionInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionInfoValidationError) ErrorName() string { return "TransactionInfoValidationError" }

// Error satisfies the builtin error interface
func (e TransactionInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactionInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionInfoValidationError{}

// Validate checks the field values on PaymentChannelInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PaymentChannelInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PaymentChannelInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PaymentChannelInfoMultiError, or nil if none found.
func (m *PaymentChannelInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *PaymentChannelInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChannelId

	// no validation rules for Type

	// no validation rules for Category

	// no validation rules for OperatorId

	// no validation rules for PaymentMethod

	// no validation rules for Tag

	// no validation rules for Name

	// no validation rules for PaymentMethodId

	// no validation rules for Currency

	// no validation rules for Country

	// no validation rules for Method

	// no validation rules for Logo

	// no validation rules for MinDepositAmount

	// no validation rules for MaxDepositAmount

	// no validation rules for MinWithdrawAmount

	// no validation rules for MaxWithdrawAmount

	// no validation rules for Eat

	if all {
		switch v := interface{}(m.GetDepositSchema()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PaymentChannelInfoValidationError{
					field:  "DepositSchema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PaymentChannelInfoValidationError{
					field:  "DepositSchema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDepositSchema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PaymentChannelInfoValidationError{
				field:  "DepositSchema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWithdrawSchema()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PaymentChannelInfoValidationError{
					field:  "WithdrawSchema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PaymentChannelInfoValidationError{
					field:  "WithdrawSchema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWithdrawSchema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PaymentChannelInfoValidationError{
				field:  "WithdrawSchema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PaymentChannelInfoMultiError(errors)
	}

	return nil
}

// PaymentChannelInfoMultiError is an error wrapping multiple validation errors
// returned by PaymentChannelInfo.ValidateAll() if the designated constraints
// aren't met.
type PaymentChannelInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaymentChannelInfoMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaymentChannelInfoMultiError) AllErrors() []error { return m }

// PaymentChannelInfoValidationError is the validation error returned by
// PaymentChannelInfo.Validate if the designated constraints aren't met.
type PaymentChannelInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentChannelInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentChannelInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentChannelInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentChannelInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentChannelInfoValidationError) ErrorName() string {
	return "PaymentChannelInfoValidationError"
}

// Error satisfies the builtin error interface
func (e PaymentChannelInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaymentChannelInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentChannelInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentChannelInfoValidationError{}

// Validate checks the field values on TransactionDetail with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TransactionDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TransactionDetail with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TransactionDetailMultiError, or nil if none found.
func (m *TransactionDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *TransactionDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTransaction()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TransactionDetailValidationError{
					field:  "Transaction",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TransactionDetailValidationError{
					field:  "Transaction",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTransaction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TransactionDetailValidationError{
				field:  "Transaction",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetChannel()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TransactionDetailValidationError{
					field:  "Channel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TransactionDetailValidationError{
					field:  "Channel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetChannel()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TransactionDetailValidationError{
				field:  "Channel",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TransactionDetailMultiError(errors)
	}

	return nil
}

// TransactionDetailMultiError is an error wrapping multiple validation errors
// returned by TransactionDetail.ValidateAll() if the designated constraints
// aren't met.
type TransactionDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransactionDetailMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransactionDetailMultiError) AllErrors() []error { return m }

// TransactionDetailValidationError is the validation error returned by
// TransactionDetail.Validate if the designated constraints aren't met.
type TransactionDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionDetailValidationError) ErrorName() string {
	return "TransactionDetailValidationError"
}

// Error satisfies the builtin error interface
func (e TransactionDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactionDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionDetailValidationError{}

// Validate checks the field values on GetPaymentTransactionByIdRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetPaymentTransactionByIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPaymentTransactionByIdRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetPaymentTransactionByIdRequestMultiError, or nil if none found.
func (m *GetPaymentTransactionByIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPaymentTransactionByIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TransactionId

	// no validation rules for Source

	if len(errors) > 0 {
		return GetPaymentTransactionByIdRequestMultiError(errors)
	}

	return nil
}

// GetPaymentTransactionByIdRequestMultiError is an error wrapping multiple
// validation errors returned by
// GetPaymentTransactionByIdRequest.ValidateAll() if the designated
// constraints aren't met.
type GetPaymentTransactionByIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPaymentTransactionByIdRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPaymentTransactionByIdRequestMultiError) AllErrors() []error { return m }

// GetPaymentTransactionByIdRequestValidationError is the validation error
// returned by GetPaymentTransactionByIdRequest.Validate if the designated
// constraints aren't met.
type GetPaymentTransactionByIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPaymentTransactionByIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPaymentTransactionByIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPaymentTransactionByIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPaymentTransactionByIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPaymentTransactionByIdRequestValidationError) ErrorName() string {
	return "GetPaymentTransactionByIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPaymentTransactionByIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPaymentTransactionByIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPaymentTransactionByIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPaymentTransactionByIdRequestValidationError{}

// Validate checks the field values on GetPaymentTransactionByIdResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetPaymentTransactionByIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPaymentTransactionByIdResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetPaymentTransactionByIdResponseMultiError, or nil if none found.
func (m *GetPaymentTransactionByIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPaymentTransactionByIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetPaymentTransactionByIdResponseValidationError{
					field:  "Detail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetPaymentTransactionByIdResponseValidationError{
					field:  "Detail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPaymentTransactionByIdResponseValidationError{
				field:  "Detail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetPaymentTransactionByIdResponseMultiError(errors)
	}

	return nil
}

// GetPaymentTransactionByIdResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetPaymentTransactionByIdResponse.ValidateAll() if the designated
// constraints aren't met.
type GetPaymentTransactionByIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPaymentTransactionByIdResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPaymentTransactionByIdResponseMultiError) AllErrors() []error { return m }

// GetPaymentTransactionByIdResponseValidationError is the validation error
// returned by GetPaymentTransactionByIdResponse.Validate if the designated
// constraints aren't met.
type GetPaymentTransactionByIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPaymentTransactionByIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPaymentTransactionByIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPaymentTransactionByIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPaymentTransactionByIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPaymentTransactionByIdResponseValidationError) ErrorName() string {
	return "GetPaymentTransactionByIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPaymentTransactionByIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPaymentTransactionByIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPaymentTransactionByIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPaymentTransactionByIdResponseValidationError{}

// Validate checks the field values on GetTransactionPageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTransactionPageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTransactionPageRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTransactionPageRequestMultiError, or nil if none found.
func (m *GetTransactionPageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTransactionPageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if m.Currency != nil {
		// no validation rules for Currency
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.StartTime != nil {

		if all {
			switch v := interface{}(m.GetStartTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTransactionPageRequestValidationError{
						field:  "StartTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTransactionPageRequestValidationError{
						field:  "StartTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTransactionPageRequestValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.EndTime != nil {

		if all {
			switch v := interface{}(m.GetEndTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTransactionPageRequestValidationError{
						field:  "EndTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTransactionPageRequestValidationError{
						field:  "EndTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEndTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTransactionPageRequestValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTransactionPageRequestMultiError(errors)
	}

	return nil
}

// GetTransactionPageRequestMultiError is an error wrapping multiple validation
// errors returned by GetTransactionPageRequest.ValidateAll() if the
// designated constraints aren't met.
type GetTransactionPageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTransactionPageRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTransactionPageRequestMultiError) AllErrors() []error { return m }

// GetTransactionPageRequestValidationError is the validation error returned by
// GetTransactionPageRequest.Validate if the designated constraints aren't met.
type GetTransactionPageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransactionPageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransactionPageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransactionPageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransactionPageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransactionPageRequestValidationError) ErrorName() string {
	return "GetTransactionPageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransactionPageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransactionPageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransactionPageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransactionPageRequestValidationError{}

// Validate checks the field values on GetTransactionPageResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTransactionPageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTransactionPageResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTransactionPageResponseMultiError, or nil if none found.
func (m *GetTransactionPageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTransactionPageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTransactions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTransactionPageResponseValidationError{
						field:  fmt.Sprintf("Transactions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTransactionPageResponseValidationError{
						field:  fmt.Sprintf("Transactions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTransactionPageResponseValidationError{
					field:  fmt.Sprintf("Transactions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalCount

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return GetTransactionPageResponseMultiError(errors)
	}

	return nil
}

// GetTransactionPageResponseMultiError is an error wrapping multiple
// validation errors returned by GetTransactionPageResponse.ValidateAll() if
// the designated constraints aren't met.
type GetTransactionPageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTransactionPageResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTransactionPageResponseMultiError) AllErrors() []error { return m }

// GetTransactionPageResponseValidationError is the validation error returned
// by GetTransactionPageResponse.Validate if the designated constraints aren't met.
type GetTransactionPageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransactionPageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransactionPageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransactionPageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransactionPageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransactionPageResponseValidationError) ErrorName() string {
	return "GetTransactionPageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransactionPageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransactionPageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransactionPageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransactionPageResponseValidationError{}
