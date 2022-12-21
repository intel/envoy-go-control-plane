// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/bumping/v3/bumping.proto

package bumpingv3

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

// Validate checks the field values on Bumping with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Bumping) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bumping with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BumpingMultiError, or nil if none found.
func (m *Bumping) ValidateAll() error {
	return m.validate(true)
}

func (m *Bumping) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := BumpingValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Cluster

	for idx, item := range m.GetAccessLog() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BumpingValidationError{
						field:  fmt.Sprintf("AccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BumpingValidationError{
						field:  fmt.Sprintf("AccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BumpingValidationError{
					field:  fmt.Sprintf("AccessLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if wrapper := m.GetMaxConnectAttempts(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			err := BumpingValidationError{
				field:  "MaxConnectAttempts",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetTlsCertificateProviderInstance()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BumpingValidationError{
					field:  "TlsCertificateProviderInstance",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BumpingValidationError{
					field:  "TlsCertificateProviderInstance",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTlsCertificateProviderInstance()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BumpingValidationError{
				field:  "TlsCertificateProviderInstance",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BumpingMultiError(errors)
	}

	return nil
}

// BumpingMultiError is an error wrapping multiple validation errors returned
// by Bumping.ValidateAll() if the designated constraints aren't met.
type BumpingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BumpingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BumpingMultiError) AllErrors() []error { return m }

// BumpingValidationError is the validation error returned by Bumping.Validate
// if the designated constraints aren't met.
type BumpingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BumpingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BumpingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BumpingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BumpingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BumpingValidationError) ErrorName() string { return "BumpingValidationError" }

// Error satisfies the builtin error interface
func (e BumpingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBumping.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BumpingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BumpingValidationError{}
