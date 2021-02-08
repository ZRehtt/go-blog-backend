package models

import (
	"time"

	"gorm.io/gorm"
)

//Model 标准模型
type Model struct {
	ID        uint            `json:"id" gorm:"primaryKey;comment:ID"`
	CreatedAt *time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt *time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
