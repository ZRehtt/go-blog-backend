package app

//状态码常量
const (
	HttpSuccess = 200
	HttpError   = 500

	//code=1000...用户模块的错误
	ErrorInvalidParams             = 1001
	ErrorNotFound                  = 1002
	UnauthorizedAuthNotExist       = 1003
	ErrorUnauthorizedToken         = 1004
	ErrorUnauthorizedTokenTimeout  = 1005
	ErrorUnauthorizedTokenGenerate = 1006
	ErrorTooManyRequests           = 1007

	//code=2000...文章模块的错误
	ErrorArticleNotFound = 2001 //文章未找到

	//code=3000...标签模块的错误
)

var codeMsg = map[int]string{
	HttpSuccess:                    "OK",
	HttpError:                      "Error",
	ErrorInvalidParams:             "请求参数错误",
	ErrorNotFound:                  "资源找不到",
	UnauthorizedAuthNotExist:       "鉴权失败，找不到对应的key和secret",
	ErrorUnauthorizedToken:         "鉴权失败，token错误",
	ErrorUnauthorizedTokenTimeout:  "鉴权失败，token超时",
	ErrorUnauthorizedTokenGenerate: "鉴权失败，token生成失败",
	ErrorTooManyRequests:           "请求次数过多",
}

func GetCodeMsg(code int) string {
	return codeMsg[code]
}
