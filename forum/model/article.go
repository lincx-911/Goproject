package model

//Article 文章模型
type Article struct {
	ID         int       `json:"id" grom:"AUTO_INCREMENT;"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserID     int       `json:"user_id"`
	Author     User      `json:"author" gorm:"foreignkey:UserID"`
	CategoryID int       `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignkey:CategoryID"`
	Likenum    int       `json:"likenum" gorm:"default:0"`
	Commentnum int       `json:"commentnum" gorm:"default:0"`
	Posttime   LocalTime `json:"posttime"`
	Updatetime LocalTime `json:"updatetime"`
	Tag        []Atag    `gorm:"many2many:article_tag;ASSOCIATION_FOREIGNKEY:ID;FOREIGNKEY:ID" json:"tag"`
}

//ArticleTag 文章标签关联
type ArticleTag struct {
	ID        int `json:"id" `
	ArticleID int `json:"article_id"`
	TagID     int `json:"tag_id"`
}
//UserFocus 关注文章
type UserFocus struct{
	UserID int `json:"user_id"`
	ArticleID int `json:"article_id"`
}

//UserLike 用户点赞的文章
type UserLike struct{
	UserID int `json:"user_id"`
	ArticleID int `json:"article_id"`
}
