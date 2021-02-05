package service

import "github.com/ZRehtt/go-blog-backend/models"

//Tag ...
type Tag struct {
	ID         uint
	Name       string
	CreatedBy  string
	UpdatedBy  string
	State      int
	PageNumber int
	PageSize   int
}

//ExistByName ...
func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

//Count ...
func (t *Tag) Count() (int64, error) {
	return models.GetTagTotal(t.getMaps())
}

//
func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = 0
	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}
	return maps
}
