package model

//Utag 文章标签
type Utag struct {
	ID       int    `json:"id"`
	Utagname string `json:"utagname"`
}

//UserTag 文章与标签关联
type UserTag struct {
	UserID int `json:"user_id"`
	UtagID int `json:"utag_id"`
}
