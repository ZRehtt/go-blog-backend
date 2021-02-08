package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User ...
type User struct {
	Model
	Username     *string `json:"username" gorm:"type:varchar(100);not null;comment:用户名"`
	PasswordHash *[]byte `json:"passwordHash" gorm:"type:varchar(100);not null;comment:用户密码"`
}

//Verify 检验必填字段
func (u *User) Verify() error {
	if u.Username == nil || (u.Username != nil && len(*u.Username) == 0) {
		return errors.New("用户名不能为空")
	}
	return nil
}

//CreateUser 新增用户
func CreateUser(user User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//CheckUserByName 根据用户名确认用户是否存在
func CheckUserByName(username string) bool {
	var user User
	err := db.Select("id").Where("username = ? AND deleted_at IS NULL", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return true
	}
	if user.ID > 0 {
		return true
	}
	return false
}

//GetUserByID 根据ID查询用户
func GetUserByID(id uint) (*User, error) {
	user := &User{}
	if err := db.Table("user").Where("id = ? AND deleted_at IS NULL", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//GetUserByName 根据Name查询用户
func GetUserByName(username string) (*User, error) {
	user := &User{}
	if err := db.Table("user").Where("username = ? AND deleted_at IS NULL", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
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
