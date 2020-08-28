package model

//Atag 文章标签
type Atag struct{
	ID int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Atagname string `json:"atagname" gorm:"unique"`
	Article []Article `gorm:"many2many:article_tag"`
}