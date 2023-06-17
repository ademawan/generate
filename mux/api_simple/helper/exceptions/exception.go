package exceptions

import (
	"errors"
)

var ErrSystem = errors.New("system error")
var ErrUnauthorized = errors.New("unauthorized")
var ErrForbidden = errors.New("forbidden")
var ErrNotFound = errors.New("error not found")
var ErrBadRequest = errors.New("bad request")
var ErrEligibility = errors.New("error not eligible")
var ErrInvalidPayload = errors.New("invalid payload")
var ErrOTPExpired = errors.New("otp expired")
var ErrRateLimit = errors.New("reach max limit request")
var ErrRateLimitSubmitOTP = errors.New("reach max limit submit otp")
var ErrConnectionMBMessage = errors.New("Failed Initializing Broker Connection")
var ErrValidDecrypt = errors.New("vailed decrypt")
var ErrRequestRegister = errors.New("Request Register Error")
var ErrDoRegister = errors.New("Do Register Error")
var ErrReadResponse = errors.New("Read Response Error")
var ErrJsonUnmarshal = errors.New("Json Unmarshal Error")
var ErrNotRegistered = errors.New("email not registered")
var ErrTomanyRequest = errors.New("max Rate Limiting Reached, Please try after some time")

func GetException(err error) error {
	return err
}
