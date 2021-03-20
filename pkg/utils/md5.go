package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5 对字符串进行md5加密
func MD5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
