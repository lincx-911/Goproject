package dao

import (
	"fmt"
	"forum/model"
	orm "forum/respository"
)

//AddCategory 添加分类
func AddCategory(category model.Category) error {
	return orm.DB.Create(&category).Error
}

//FindCategoryByname 通过名称查找分类
func FindCategoryByname(name string) (cate model.Category,ok bool){
	ok=false
	ok = orm.DB.Where("category_name = ?",name).First(&cate).RecordNotFound()
	fmt.Println("catelast",cate)
	return
}