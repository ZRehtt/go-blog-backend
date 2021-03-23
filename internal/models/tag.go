package models

//Tag ...
type Tag struct {
	*Model
	Name  string `json:"name"`
	State int8   `json:"state"`
}
