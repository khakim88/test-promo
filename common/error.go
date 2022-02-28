package common

// Error is a struct of error obj
// code is for error code
// status code is the first 3 digits of the error code
type Error struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) CodeErr() int {
	return e.Code
}

func (e *Error) DataErr() map[string]interface{} {
	return e.Data
}

func (e *Error) WithData(data map[string]interface{}) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Data:    data,
	}
}

func NewLegacyErrorMessage(code int, message string, data map[string]interface{}) *Error {
	return &Error{Code: code, Message: message, Data: data}
}

var (
	ErrInternal     = &Error{50000, "Sorry. Unable to complete your request", nil}
	ErrTokenExpired = &Error{50001, "Sorry, token is expired", nil}
	ErrTokenInvalid = &Error{50002, "Soory, token is invalid", nil}
)
