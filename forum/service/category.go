package service

import (
	"errors"
	"forum/dao"
	"forum/model"
)

//AddCategoryService 添加分类
func AddCategoryService(category model.Category) (bool,error) {
	if _,notfind:=dao.FindCategoryByname(category.CategoryName);!notfind{
		return true,errors.New("Had Exited!")
	}
	if err:=dao.AddCategory(category);err!=nil{
		return false,err
	}
	return true,nil
}