package models

//Article ...
type Article struct {
	*Model
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Content  string `json:"content"`
	CoverURL string `json:"coverURL"`
	State    int8   `json:"state"`
}
