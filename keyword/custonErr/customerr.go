package custonErr

import "fmt"

// CError 通用的错误处理类
type CError struct {
	code int
	msg  string
}

// New 创建一个错误对象
func New(code int, msg string) *CError {
	return &CError{code: code, msg: msg}
}

// Data 返回code和msg内容
func (e *CError) Data() (code int, msg string) {
	code, msg = e.code, e.msg
	return
}

// Error  实现error interface的接口
func (e *CError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.code, e.msg)
}

// Error  实现error interface的接口
func (e *CError) ErrorMsg() string {
	return e.msg
}

// Code 返回错误码
func (e *CError) Code() int {
	return e.code
}
