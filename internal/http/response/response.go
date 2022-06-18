package response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ErrorCode string

const (
	InvalidParameter ErrorCode = "InvalidParameter"
	Unauthenticated  ErrorCode = "Unauthenticated"
	SystemError      ErrorCode = "SystemError"
)

type Response struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Status  int       `json:"-"`
}

func (r *Response) Error() string {
	return r.Message
}

func NewInvalidParameter(message interface{}) *Response {
	return &Response{
		Code:    InvalidParameter,
		Message: fmt.Sprint(message),
		Status:  fiber.StatusInternalServerError,
	}
}

func NewUnauthenticated(message interface{}) *Response {
	return &Response{
		Code:    Unauthenticated,
		Message: fmt.Sprint(message),
		Status:  fiber.StatusUnauthorized,
	}
}

func NewSystemError(message interface{}) *Response {
	return &Response{
		Code:    SystemError,
		Message: fmt.Sprint(message),
		Status:  fiber.StatusInternalServerError,
	}
}
