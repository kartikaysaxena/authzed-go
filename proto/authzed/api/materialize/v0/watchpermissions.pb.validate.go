// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authzed/api/materialize/v0/watchpermissions.proto

package v0

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

// Validate checks the field values on WatchPermissionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WatchPermissionsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WatchPermissionsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WatchPermissionsRequestMultiError, or nil if none found.
func (m *WatchPermissionsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WatchPermissionsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPermissions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WatchPermissionsRequestValidationError{
						field:  fmt.Sprintf("Permissions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WatchPermissionsRequestValidationError{
						field:  fmt.Sprintf("Permissions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WatchPermissionsRequestValidationError{
					field:  fmt.Sprintf("Permissions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetOptionalStartingAfter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WatchPermissionsRequestValidationError{
					field:  "OptionalStartingAfter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WatchPermissionsRequestValidationError{
					field:  "OptionalStartingAfter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOptionalStartingAfter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WatchPermissionsRequestValidationError{
				field:  "OptionalStartingAfter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return WatchPermissionsRequestMultiError(errors)
	}

	return nil
}

// WatchPermissionsRequestMultiError is an error wrapping multiple validation
// errors returned by WatchPermissionsRequest.ValidateAll() if the designated
// constraints aren't met.
type WatchPermissionsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WatchPermissionsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WatchPermissionsRequestMultiError) AllErrors() []error { return m }

// WatchPermissionsRequestValidationError is the validation error returned by
// WatchPermissionsRequest.Validate if the designated constraints aren't met.
type WatchPermissionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WatchPermissionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WatchPermissionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WatchPermissionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WatchPermissionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WatchPermissionsRequestValidationError) ErrorName() string {
	return "WatchPermissionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e WatchPermissionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchPermissionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WatchPermissionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WatchPermissionsRequestValidationError{}

// Validate checks the field values on WatchedPermission with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WatchedPermission) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WatchedPermission with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WatchedPermissionMultiError, or nil if none found.
func (m *WatchedPermission) ValidateAll() error {
	return m.validate(true)
}

func (m *WatchedPermission) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ResourceType

	// no validation rules for Permission

	// no validation rules for SubjectType

	// no validation rules for OptionalSubjectRelation

	if len(errors) > 0 {
		return WatchedPermissionMultiError(errors)
	}

	return nil
}

// WatchedPermissionMultiError is an error wrapping multiple validation errors
// returned by WatchedPermission.ValidateAll() if the designated constraints
// aren't met.
type WatchedPermissionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WatchedPermissionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WatchedPermissionMultiError) AllErrors() []error { return m }

// WatchedPermissionValidationError is the validation error returned by
// WatchedPermission.Validate if the designated constraints aren't met.
type WatchedPermissionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WatchedPermissionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WatchedPermissionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WatchedPermissionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WatchedPermissionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WatchedPermissionValidationError) ErrorName() string {
	return "WatchedPermissionValidationError"
}

// Error satisfies the builtin error interface
func (e WatchedPermissionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchedPermission.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WatchedPermissionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WatchedPermissionValidationError{}

// Validate checks the field values on WatchPermissionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WatchPermissionsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WatchPermissionsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WatchPermissionsResponseMultiError, or nil if none found.
func (m *WatchPermissionsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *WatchPermissionsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Response.(type) {
	case *WatchPermissionsResponse_Change:
		if v == nil {
			err := WatchPermissionsResponseValidationError{
				field:  "Response",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetChange()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WatchPermissionsResponseValidationError{
						field:  "Change",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WatchPermissionsResponseValidationError{
						field:  "Change",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetChange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WatchPermissionsResponseValidationError{
					field:  "Change",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *WatchPermissionsResponse_CompletedRevision:
		if v == nil {
			err := WatchPermissionsResponseValidationError{
				field:  "Response",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetCompletedRevision()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WatchPermissionsResponseValidationError{
						field:  "CompletedRevision",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WatchPermissionsResponseValidationError{
						field:  "CompletedRevision",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCompletedRevision()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WatchPermissionsResponseValidationError{
					field:  "CompletedRevision",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return WatchPermissionsResponseMultiError(errors)
	}

	return nil
}

// WatchPermissionsResponseMultiError is an error wrapping multiple validation
// errors returned by WatchPermissionsResponse.ValidateAll() if the designated
// constraints aren't met.
type WatchPermissionsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WatchPermissionsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WatchPermissionsResponseMultiError) AllErrors() []error { return m }

// WatchPermissionsResponseValidationError is the validation error returned by
// WatchPermissionsResponse.Validate if the designated constraints aren't met.
type WatchPermissionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WatchPermissionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WatchPermissionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WatchPermissionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WatchPermissionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WatchPermissionsResponseValidationError) ErrorName() string {
	return "WatchPermissionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e WatchPermissionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchPermissionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WatchPermissionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WatchPermissionsResponseValidationError{}

// Validate checks the field values on PermissionChange with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PermissionChange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PermissionChange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PermissionChangeMultiError, or nil if none found.
func (m *PermissionChange) ValidateAll() error {
	return m.validate(true)
}

func (m *PermissionChange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRevision()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PermissionChangeValidationError{
					field:  "Revision",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PermissionChangeValidationError{
					field:  "Revision",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PermissionChangeValidationError{
				field:  "Revision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PermissionChangeValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PermissionChangeValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PermissionChangeValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Permission

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PermissionChangeValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PermissionChangeValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PermissionChangeValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Permissionship

	if len(errors) > 0 {
		return PermissionChangeMultiError(errors)
	}

	return nil
}

// PermissionChangeMultiError is an error wrapping multiple validation errors
// returned by PermissionChange.ValidateAll() if the designated constraints
// aren't met.
type PermissionChangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PermissionChangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PermissionChangeMultiError) AllErrors() []error { return m }

// PermissionChangeValidationError is the validation error returned by
// PermissionChange.Validate if the designated constraints aren't met.
type PermissionChangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PermissionChangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PermissionChangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PermissionChangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PermissionChangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PermissionChangeValidationError) ErrorName() string { return "PermissionChangeValidationError" }

// Error satisfies the builtin error interface
func (e PermissionChangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPermissionChange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PermissionChangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PermissionChangeValidationError{}