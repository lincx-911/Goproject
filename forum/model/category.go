package model

//Category 文章分类
type Category struct{
	ID int `json:"id" grom:"AUTO_INCREMENT"`
	CategoryName string `json:"category_name"`
}