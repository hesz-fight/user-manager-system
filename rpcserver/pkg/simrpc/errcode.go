package simrpc

import (
	"fmt"
)

var (
	ErrorMethodHasExist = NewError(1000001, "方法已存在")
	ErrorMehodNotFound  = NewError(1000002, "方法未找到")
)

type CommonError struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) *CommonError {
	return &CommonError{Code: code, Msg: msg}
}

func (c *CommonError) Error() string {
	return fmt.Sprintf("%d:%s", c.Code, c.Msg)
}

func (c *CommonError) GetCode() int {
	return c.Code
}

func (c *CommonError) GetMsg() string {
	return c.Msg
}
