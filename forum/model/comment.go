package model

//Comment 评论
type Comment struct{
	ID int `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	Raid int `json:"raid"`
	Paid int `json:"paid"`
	Content string `json:"content"`
	UserID  int `json:"user_id"`
	Author User `json:"author" gorm:"foreignkey:UserID"`
	TouserID int `json:"touser_id"`
	Touser User `json:"touser" gorm:"foreignkey:TouserID"`
	Likenum int `json:"likenum" gorm:"default:0"`
	Commentnum int `json:"commentnum" gorm:"default:0"`
	PostTime LocalTime `json:"post_time"`
}
//UserClike 用户评论点赞
type UserClike struct{
	UserID int `json:"user_id"`
	CommentID int `json:"comment_id"`
}