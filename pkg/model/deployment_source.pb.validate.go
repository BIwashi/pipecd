// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/model/deployment_source.proto

package model

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

// Validate checks the field values on DeploymentSource with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeploymentSource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeploymentSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeploymentSourceMultiError, or nil if none found.
func (m *DeploymentSource) ValidateAll() error {
	return m.validate(true)
}

func (m *DeploymentSource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ApplicationDirectory

	// no validation rules for Revision

	// no validation rules for ApplicationConfig

	// no validation rules for ApplicationConfigFilename

	if len(errors) > 0 {
		return DeploymentSourceMultiError(errors)
	}

	return nil
}

// DeploymentSourceMultiError is an error wrapping multiple validation errors
// returned by DeploymentSource.ValidateAll() if the designated constraints
// aren't met.
type DeploymentSourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentSourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentSourceMultiError) AllErrors() []error { return m }

// DeploymentSourceValidationError is the validation error returned by
// DeploymentSource.Validate if the designated constraints aren't met.
type DeploymentSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentSourceValidationError) ErrorName() string { return "DeploymentSourceValidationError" }

// Error satisfies the builtin error interface
func (e DeploymentSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeploymentSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentSourceValidationError{}
