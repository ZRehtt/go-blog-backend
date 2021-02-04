package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	Model
	Username     *string `json:"username" gorm:"username;comment:用户名"`
	PasswordHash *[]byte `json:"-" gorm:"password_hash;comment:用户密码"`
}

//Verify 检验必填字段
func (u *User) Verify() error {
	if u.Username == nil || (u.Username != nil && len(*u.Username) == 0) {
		return errors.New("用户名不能为空")
	}
	return nil
}

//SetPassword 将明文密码加密
func (u *User) SetPassword(password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.PasswordHash = &hash
	return nil
}

//CheckPassword ...
func (u *User) CheckPassword(password string) error {
	//CompareHashAndPassword比较bcrypt哈希密码和其可能的明文等价成本。成功时返回nil，失败时返回错误。
	if u.PasswordHash != nil && len(*u.PasswordHash) == 0 {
		return errors.New("密码未设置")
	}
	return bcrypt.CompareHashAndPassword(*u.PasswordHash, []byte(password))
}

//HashPassword 将密码hash加密
func HashPassword(password string) ([]byte, error) {
	//GenerateFromPassword 以给定的成本返回密码的 bcrypt 哈希值。
	//如果给定的成本小于 MinCost，成本将被设置为 DefaultCost，也就是10。
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
