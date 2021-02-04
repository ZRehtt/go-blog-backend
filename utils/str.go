package utils

import (
	"math/rand"
	"strings"
	"time"
)

//RandStr 生成随机字符串
func RandStr(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	var strB strings.Builder
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		c := str[r.Intn(len(str))]
		strB.WriteByte(c)
	}
	return strB.String()
}
