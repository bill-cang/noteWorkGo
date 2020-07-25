// Package terror 通用的错误类
// Copyright 2017 Aladinfun
// Author yison@shugadating.com
// The software is NOT free software
//Any terms of copy or usage of the code should be under the permission of the author

package terror

import (
	"fmt"
)

// TError 通用的错误处理类
type TError struct {
	code int
	msg  string
}

// New 创建一个错误对象
func New(code ERRORCODE, msg string) *TError {
	return &TError{code: int(code), msg: msg}
}

// Data 返回code和msg内容
func (e *TError) Data() (code int, msg string) {
	code, msg = e.code, e.msg
	return
}

// Error  实现error interface的接口
func (e *TError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.code, e.msg)
}

// Error  实现error interface的接口
func (e *TError) ErrorMsg() string {
	return e.msg
}

// Code 返回错误码
func (e *TError) Code() int {
	return e.code
}
