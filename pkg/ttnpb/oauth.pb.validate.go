// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _oauth_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on
// OAuthClientAuthorizationIdentifiers with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *OAuthClientAuthorizationIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OAuthClientAuthorizationIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "user_ids":

			if m.GetUserIds() == nil {
				return OAuthClientAuthorizationIdentifiersValidationError{
					field:  "user_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthClientAuthorizationIdentifiersValidationError{
						field:  "user_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "client_ids":

			if m.GetClientIds() == nil {
				return OAuthClientAuthorizationIdentifiersValidationError{
					field:  "client_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetClientIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthClientAuthorizationIdentifiersValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return OAuthClientAuthorizationIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OAuthClientAuthorizationIdentifiersValidationError is the validation error
// returned by OAuthClientAuthorizationIdentifiers.ValidateFields if the
// designated constraints aren't met.
type OAuthClientAuthorizationIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuthClientAuthorizationIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuthClientAuthorizationIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuthClientAuthorizationIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuthClientAuthorizationIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuthClientAuthorizationIdentifiersValidationError) ErrorName() string {
	return "OAuthClientAuthorizationIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e OAuthClientAuthorizationIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuthClientAuthorizationIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuthClientAuthorizationIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuthClientAuthorizationIdentifiersValidationError{}

// ValidateFields checks the field values on OAuthClientAuthorization with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OAuthClientAuthorization) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OAuthClientAuthorizationFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "user_ids":

			if m.GetUserIds() == nil {
				return OAuthClientAuthorizationValidationError{
					field:  "user_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthClientAuthorizationValidationError{
						field:  "user_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "client_ids":

			if m.GetClientIds() == nil {
				return OAuthClientAuthorizationValidationError{
					field:  "client_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetClientIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthClientAuthorizationValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "rights":

		case "created_at":

			if v, ok := interface{}(m.GetCreatedAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthClientAuthorizationValidationError{
						field:  "created_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "updated_at":

			if v, ok := interface{}(m.GetUpdatedAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthClientAuthorizationValidationError{
						field:  "updated_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return OAuthClientAuthorizationValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OAuthClientAuthorizationValidationError is the validation error returned by
// OAuthClientAuthorization.ValidateFields if the designated constraints
// aren't met.
type OAuthClientAuthorizationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuthClientAuthorizationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuthClientAuthorizationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuthClientAuthorizationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuthClientAuthorizationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuthClientAuthorizationValidationError) ErrorName() string {
	return "OAuthClientAuthorizationValidationError"
}

// Error satisfies the builtin error interface
func (e OAuthClientAuthorizationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuthClientAuthorization.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuthClientAuthorizationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuthClientAuthorizationValidationError{}

// ValidateFields checks the field values on OAuthClientAuthorizations with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OAuthClientAuthorizations) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OAuthClientAuthorizationsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "authorizations":

			for idx, item := range m.GetAuthorizations() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return OAuthClientAuthorizationsValidationError{
							field:  fmt.Sprintf("authorizations[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return OAuthClientAuthorizationsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OAuthClientAuthorizationsValidationError is the validation error returned by
// OAuthClientAuthorizations.ValidateFields if the designated constraints
// aren't met.
type OAuthClientAuthorizationsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuthClientAuthorizationsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuthClientAuthorizationsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuthClientAuthorizationsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuthClientAuthorizationsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuthClientAuthorizationsValidationError) ErrorName() string {
	return "OAuthClientAuthorizationsValidationError"
}

// Error satisfies the builtin error interface
func (e OAuthClientAuthorizationsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuthClientAuthorizations.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuthClientAuthorizationsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuthClientAuthorizationsValidationError{}

// ValidateFields checks the field values on
// ListOAuthClientAuthorizationsRequest with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *ListOAuthClientAuthorizationsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ListOAuthClientAuthorizationsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "user_ids":

			if m.GetUserIds() == nil {
				return ListOAuthClientAuthorizationsRequestValidationError{
					field:  "user_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListOAuthClientAuthorizationsRequestValidationError{
						field:  "user_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":

			if _, ok := _ListOAuthClientAuthorizationsRequest_Order_InLookup[m.GetOrder()]; !ok {
				return ListOAuthClientAuthorizationsRequestValidationError{
					field:  "order",
					reason: "value must be in list [ created_at -created_at]",
				}
			}

		case "limit":

			if m.GetLimit() > 1000 {
				return ListOAuthClientAuthorizationsRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		default:
			return ListOAuthClientAuthorizationsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ListOAuthClientAuthorizationsRequestValidationError is the validation error
// returned by ListOAuthClientAuthorizationsRequest.ValidateFields if the
// designated constraints aren't met.
type ListOAuthClientAuthorizationsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOAuthClientAuthorizationsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOAuthClientAuthorizationsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOAuthClientAuthorizationsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOAuthClientAuthorizationsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOAuthClientAuthorizationsRequestValidationError) ErrorName() string {
	return "ListOAuthClientAuthorizationsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListOAuthClientAuthorizationsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOAuthClientAuthorizationsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOAuthClientAuthorizationsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOAuthClientAuthorizationsRequestValidationError{}

var _ListOAuthClientAuthorizationsRequest_Order_InLookup = map[string]struct{}{
	"":            {},
	"created_at":  {},
	"-created_at": {},
}

// ValidateFields checks the field values on OAuthAuthorizationCode with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OAuthAuthorizationCode) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OAuthAuthorizationCodeFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "user_ids":

			if m.GetUserIds() == nil {
				return OAuthAuthorizationCodeValidationError{
					field:  "user_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAuthorizationCodeValidationError{
						field:  "user_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "user_session_id":

			if utf8.RuneCountInString(m.GetUserSessionId()) > 64 {
				return OAuthAuthorizationCodeValidationError{
					field:  "user_session_id",
					reason: "value length must be at most 64 runes",
				}
			}

		case "client_ids":

			if m.GetClientIds() == nil {
				return OAuthAuthorizationCodeValidationError{
					field:  "client_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetClientIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAuthorizationCodeValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "rights":

		case "code":
			// no validation rules for Code
		case "redirect_uri":

			if _, err := url.Parse(m.GetRedirectUri()); err != nil {
				return OAuthAuthorizationCodeValidationError{
					field:  "redirect_uri",
					reason: "value must be a valid URI",
					cause:  err,
				}
			}

		case "state":
			// no validation rules for State
		case "created_at":

			if v, ok := interface{}(m.GetCreatedAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAuthorizationCodeValidationError{
						field:  "created_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "expires_at":

			if v, ok := interface{}(m.GetExpiresAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAuthorizationCodeValidationError{
						field:  "expires_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return OAuthAuthorizationCodeValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OAuthAuthorizationCodeValidationError is the validation error returned by
// OAuthAuthorizationCode.ValidateFields if the designated constraints aren't met.
type OAuthAuthorizationCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuthAuthorizationCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuthAuthorizationCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuthAuthorizationCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuthAuthorizationCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuthAuthorizationCodeValidationError) ErrorName() string {
	return "OAuthAuthorizationCodeValidationError"
}

// Error satisfies the builtin error interface
func (e OAuthAuthorizationCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuthAuthorizationCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuthAuthorizationCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuthAuthorizationCodeValidationError{}

// ValidateFields checks the field values on OAuthAccessTokenIdentifiers with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *OAuthAccessTokenIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OAuthAccessTokenIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "user_ids":

			if m.GetUserIds() == nil {
				return OAuthAccessTokenIdentifiersValidationError{
					field:  "user_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAccessTokenIdentifiersValidationError{
						field:  "user_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "client_ids":

			if m.GetClientIds() == nil {
				return OAuthAccessTokenIdentifiersValidationError{
					field:  "client_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetClientIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAccessTokenIdentifiersValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "id":
			// no validation rules for Id
		default:
			return OAuthAccessTokenIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OAuthAccessTokenIdentifiersValidationError is the validation error returned
// by OAuthAccessTokenIdentifiers.ValidateFields if the designated constraints
// aren't met.
type OAuthAccessTokenIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuthAccessTokenIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuthAccessTokenIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuthAccessTokenIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuthAccessTokenIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuthAccessTokenIdentifiersValidationError) ErrorName() string {
	return "OAuthAccessTokenIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e OAuthAccessTokenIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuthAccessTokenIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuthAccessTokenIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuthAccessTokenIdentifiersValidationError{}

// ValidateFields checks the field values on OAuthAccessToken with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OAuthAccessToken) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OAuthAccessTokenFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "user_ids":

			if m.GetUserIds() == nil {
				return OAuthAccessTokenValidationError{
					field:  "user_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAccessTokenValidationError{
						field:  "user_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "user_session_id":

			if utf8.RuneCountInString(m.GetUserSessionId()) > 64 {
				return OAuthAccessTokenValidationError{
					field:  "user_session_id",
					reason: "value length must be at most 64 runes",
				}
			}

		case "client_ids":

			if m.GetClientIds() == nil {
				return OAuthAccessTokenValidationError{
					field:  "client_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetClientIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAccessTokenValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "id":
			// no validation rules for Id
		case "access_token":
			// no validation rules for AccessToken
		case "refresh_token":
			// no validation rules for RefreshToken
		case "rights":

		case "created_at":

			if v, ok := interface{}(m.GetCreatedAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAccessTokenValidationError{
						field:  "created_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "expires_at":

			if v, ok := interface{}(m.GetExpiresAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAccessTokenValidationError{
						field:  "expires_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "session_expires_at":

			if v, ok := interface{}(m.GetSessionExpiresAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return OAuthAccessTokenValidationError{
						field:  "session_expires_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return OAuthAccessTokenValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OAuthAccessTokenValidationError is the validation error returned by
// OAuthAccessToken.ValidateFields if the designated constraints aren't met.
type OAuthAccessTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuthAccessTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuthAccessTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuthAccessTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuthAccessTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuthAccessTokenValidationError) ErrorName() string { return "OAuthAccessTokenValidationError" }

// Error satisfies the builtin error interface
func (e OAuthAccessTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuthAccessToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuthAccessTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuthAccessTokenValidationError{}

// ValidateFields checks the field values on OAuthAccessTokens with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OAuthAccessTokens) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OAuthAccessTokensFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "tokens":

			for idx, item := range m.GetTokens() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return OAuthAccessTokensValidationError{
							field:  fmt.Sprintf("tokens[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return OAuthAccessTokensValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OAuthAccessTokensValidationError is the validation error returned by
// OAuthAccessTokens.ValidateFields if the designated constraints aren't met.
type OAuthAccessTokensValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuthAccessTokensValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuthAccessTokensValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuthAccessTokensValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuthAccessTokensValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuthAccessTokensValidationError) ErrorName() string {
	return "OAuthAccessTokensValidationError"
}

// Error satisfies the builtin error interface
func (e OAuthAccessTokensValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuthAccessTokens.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuthAccessTokensValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuthAccessTokensValidationError{}

// ValidateFields checks the field values on ListOAuthAccessTokensRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ListOAuthAccessTokensRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ListOAuthAccessTokensRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "user_ids":

			if m.GetUserIds() == nil {
				return ListOAuthAccessTokensRequestValidationError{
					field:  "user_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListOAuthAccessTokensRequestValidationError{
						field:  "user_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "client_ids":

			if m.GetClientIds() == nil {
				return ListOAuthAccessTokensRequestValidationError{
					field:  "client_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetClientIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListOAuthAccessTokensRequestValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":

			if _, ok := _ListOAuthAccessTokensRequest_Order_InLookup[m.GetOrder()]; !ok {
				return ListOAuthAccessTokensRequestValidationError{
					field:  "order",
					reason: "value must be in list [ created_at -created_at]",
				}
			}

		case "limit":

			if m.GetLimit() > 1000 {
				return ListOAuthAccessTokensRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		default:
			return ListOAuthAccessTokensRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ListOAuthAccessTokensRequestValidationError is the validation error returned
// by ListOAuthAccessTokensRequest.ValidateFields if the designated
// constraints aren't met.
type ListOAuthAccessTokensRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOAuthAccessTokensRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOAuthAccessTokensRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOAuthAccessTokensRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOAuthAccessTokensRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOAuthAccessTokensRequestValidationError) ErrorName() string {
	return "ListOAuthAccessTokensRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListOAuthAccessTokensRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOAuthAccessTokensRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOAuthAccessTokensRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOAuthAccessTokensRequestValidationError{}

var _ListOAuthAccessTokensRequest_Order_InLookup = map[string]struct{}{
	"":            {},
	"created_at":  {},
	"-created_at": {},
}
