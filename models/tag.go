package models

//Tag 文章标签管理
type Tag struct {
	Model
	Name       string `json:"name,omitempty" gorm:"name;comment:标签名称"`
	CreatedBy  string `json:"createdBy,omitempty" gorm:"created_by;comment:标签创建者"`
	ModifiedBy string `json:"modifiedBy,omitempty" gorm:"modified_by;comment:标签修改者"`
	State      int    `json:"state,omitempty" gorm:"state;default:0;comment:标签状态，状态 0为禁用、1为启用"`
}
