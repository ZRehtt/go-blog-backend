package utils

import "strconv"

/**
convert 工具用于类型转换
 */
type StrTo string

//String ...
func (s StrTo) String() string {
	return string(s)
}
//Int 用于将string类型转换为int
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) Uint32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}
