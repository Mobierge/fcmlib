package fcmlib

import "fmt"

// Error
type fcmError struct {
	code    fcmErrorCode
	message string
}

func (e *fcmError) Error() string {
	return fmt.Sprintf("%s: %s", e.code.String(), e.message)
}

func (e *fcmError) ShouldRetry() bool {
	return e.code == ErrorConnection || e.code == ErrorServiceUnavailable
}

func (e *fcmError) Code() fcmErrorCode {
	return e.code
}

func newError(errorCode fcmErrorCode, message string) *fcmError {
	return &fcmError{code: errorCode, message: message}
}

type fcmErrorCode int

// Errors
const (
	ErrorUnknown               fcmErrorCode = 1
	ErrorBadRequest            fcmErrorCode = 400
	ErrorAuthentication        fcmErrorCode = 401
	ErrorRequestEntityTooLarge fcmErrorCode = 413
	ErrorServiceUnavailable    fcmErrorCode = 500
	ErrorResponseParse         fcmErrorCode = 1000
	ErrorConnection            fcmErrorCode = 1001
)

func (t fcmErrorCode) String() string {
	switch t {
	case ErrorBadRequest:
		return "bad request"
	case ErrorAuthentication:
		return "authentication error"
	case ErrorRequestEntityTooLarge:
		return "request entity too large"
	case ErrorServiceUnavailable:
		return "fcm service unavailable"
	case ErrorResponseParse:
		return "response body is not well-formed json"
	case ErrorConnection:
		return "connection error"
	default:
		return "unknown error"
	}
}

type resultError string

// Result-level error codes can be found in response.Results for each individual
// recipient.
// See https://developers.google.com/cloud-messaging/server-ref#error-codes
const (
	ResultErrorMissingRegistration       resultError = "MissingRegistration"
	ResultErrorInvalidRegistration       resultError = "InvalidRegistration"
	ResultErrorNotRegistered             resultError = "NotRegistered"
	ResultErrorMessageTooBig             resultError = "MessageTooBig"
	ResultErrorInvalidDataKey            resultError = "InvalidDataKey"
	ResultErrorInvalidTTL                resultError = "InvalidTtl"
	ResultErrorDeviceMessageRateExceeded resultError = "DeviceMessageRateExceeded"
	ResultErrorTopicsMessageRateExceeded resultError = "TopicsMessageRateExceeded"
	ResultErrorMismatchSenderID          resultError = "MismatchSenderId"
	ResultErrorInvalidPackageName        resultError = "InvalidPackageName"
	ResultErrorInternalServerError       resultError = "InternalServerError"
	ResultErrorUnavailable               resultError = "Unavailable"
)
