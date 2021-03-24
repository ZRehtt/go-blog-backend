package errcode

import (
	"fmt"
	"net/http"
)

//Error 错误信息的格式结构
type Error struct {
	//错误码
	code int
	//错误信息
	msg string
	//详细信息
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code(), e.Msg())
}

func (e *Error) MsgF(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

//错误码和状态码的转换
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case ErrorInvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case ErrorUnauthorizedToken.Code():
		fallthrough
	case ErrorUnauthorizedTokenGenerate.Code():
		fallthrough
	case ErrorUnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case ErrorTooManyRequests.Code():
		return http.StatusTooManyRequests
	case ErrorNotFound.Code():
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
