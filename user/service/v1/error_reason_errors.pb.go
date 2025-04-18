// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNSPECIFIED.String() && e.Code == 500
}

func ErrorUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUserAuthNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_AUTH_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserAuthNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_AUTH_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUserAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_ALREADY_EXISTS.String() && e.Code == 409
}

func ErrorUserAlreadyExists(format string, args ...interface{}) *errors.Error {
	return errors.New(409, ErrorReason_USER_ALREADY_EXISTS.String(), fmt.Sprintf(format, args...))
}

func IsRevokeTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_REVOKE_TOKEN_FAILED.String() && e.Code == 500
}

func ErrorRevokeTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_REVOKE_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGenerateRefreshTokenIdFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GENERATE_REFRESH_TOKEN_ID_FAILED.String() && e.Code == 500
}

func ErrorGenerateRefreshTokenIdFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GENERATE_REFRESH_TOKEN_ID_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGenerateTokenIdFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GENERATE_TOKEN_ID_FAILED.String() && e.Code == 500
}

func ErrorGenerateTokenIdFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GENERATE_TOKEN_ID_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSaveTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SAVE_TOKEN_FAILED.String() && e.Code == 500
}

func ErrorSaveTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_SAVE_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetUserAuthFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_USER_AUTH_FAILED.String() && e.Code == 500
}

func ErrorGetUserAuthFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GET_USER_AUTH_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetUserFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_USER_FAILED.String() && e.Code == 500
}

func ErrorGetUserFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GET_USER_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsUpdateUserFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UPDATE_USER_FAILED.String() && e.Code == 500
}

func ErrorUpdateUserFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_UPDATE_USER_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsLockUserAuthFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_LOCK_USER_AUTH_FAILED.String() && e.Code == 500
}

func ErrorLockUserAuthFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_LOCK_USER_AUTH_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsHashUserPasswordFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_HASH_USER_PASSWORD_FAILED.String() && e.Code == 500
}

func ErrorHashUserPasswordFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_HASH_USER_PASSWORD_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGenerateUserIdFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GENERATE_USER_ID_FAILED.String() && e.Code == 500
}

func ErrorGenerateUserIdFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GENERATE_USER_ID_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateUserFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CREATE_USER_FAILED.String() && e.Code == 500
}

func ErrorCreateUserFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_CREATE_USER_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateUserAuthFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CREATE_USER_AUTH_FAILED.String() && e.Code == 500
}

func ErrorCreateUserAuthFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_CREATE_USER_AUTH_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsUserDisabled(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_DISABLED.String() && e.Code == 401
}

func ErrorUserDisabled(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_USER_DISABLED.String(), fmt.Sprintf(format, args...))
}

func IsUserLoginBanned(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_LOGIN_BANNED.String() && e.Code == 401
}

func ErrorUserLoginBanned(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_USER_LOGIN_BANNED.String(), fmt.Sprintf(format, args...))
}

func IsInvalidUserPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_USER_PASSWORD.String() && e.Code == 401
}

func ErrorInvalidUserPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_INVALID_USER_PASSWORD.String(), fmt.Sprintf(format, args...))
}

func IsGetTokenWithRefreshTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_TOKEN_WITH_REFRESH_TOKEN_FAILED.String() && e.Code == 500
}

func ErrorGetTokenWithRefreshTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GET_TOKEN_WITH_REFRESH_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetUnexpiredTokenForUserFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_UNEXPIRED_TOKEN_FOR_USER_FAILED.String() && e.Code == 500
}

func ErrorGetUnexpiredTokenForUserFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GET_UNEXPIRED_TOKEN_FOR_USER_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsOauthProviderNotSupported(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_OAUTH_PROVIDER_NOT_SUPPORTED.String() && e.Code == 500
}

func ErrorOauthProviderNotSupported(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_OAUTH_PROVIDER_NOT_SUPPORTED.String(), fmt.Sprintf(format, args...))
}
