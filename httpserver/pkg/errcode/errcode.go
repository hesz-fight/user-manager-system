package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string
}

var (
	codes = map[int]string{}
)

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("err code exists yet: %d", code))
	}

	codes[code] = msg

	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("err code: %d, err msg: %s", e.code, e.msg)
}

func (e *Error) GetCode() int {
	return e.code
}

func (e *Error) GetDetails() []string {
	return e.details
}

func (e *Error) GetMsg() string {
	return e.msg
}

func (e *Error) WithDetils(details ...string) *Error {
	ne := *e
	ne.details = []string{}
	for _, d := range details {
		ne.details = append(ne.details, d)
	}

	return &ne
}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParams.code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.code:
		fallthrough
	case UnauthorizedTokenError.code:
		fallthrough
	case UnauthorizedTokenGenerate.code:
		fallthrough
	case UnauthorizedTokenTimeout.code:
		return http.StatusUnauthorized
	case TooManyRequests.code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
