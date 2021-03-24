package errcode

//公共错误码和信息
var (
	Success                        = NewError(0, "成功")
	ServerError                    = NewError(10000, "服务内部错误")
	ErrorInvalidParams             = NewError(10001, "参数错误")
	ErrorNotFound                  = NewError(10002, "资源找不到")
	UnauthorizedAuthNotExist       = NewError(10003, "鉴权失败，找不到对应的key和secret")
	ErrorUnauthorizedToken         = NewError(10004, "鉴权失败，token错误")
	ErrorUnauthorizedTokenTimeout  = NewError(10005, "鉴权失败，token超时")
	ErrorUnauthorizedTokenGenerate = NewError(10006, "鉴权失败，token生成失败")
	ErrorTooManyRequests           = NewError(10007, "请求次数过多")
)
