package common

import "fmt"

type LegacyError struct {
	Err string `json:"err"`
}

func NewLegacyError(errMsg string) LegacyError {
	return LegacyError{Err: errMsg}
}

func NewLegacyErrorf(errMsgFormat string, args ...interface{}) LegacyError {
	errMsg := fmt.Sprintf(errMsgFormat, args...)
	return NewLegacyError(errMsg)
}

func (e LegacyError) Error() string {
	return e.Err
}
