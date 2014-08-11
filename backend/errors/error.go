package errors

import (
	"fmt"
)

type ServerError struct {
	Err        error  `json:"error"`
	ErrMessage string `json:"errormessage"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

func (e ServerError) Error() string {
	return fmt.Sprintf("%#v", e)
}

func New(err error, message string, statusCode int) *ServerError {
	if err != nil {
		return &ServerError{
			Err:        err,
			ErrMessage: err.Error(),
			Message:    message,
			StatusCode: statusCode,
		}
	} else {
		return &ServerError{
			Err:        nil,
			ErrMessage: "",
			Message:    message,
			StatusCode: statusCode,
		}
	}
}
