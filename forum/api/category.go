package api

import (
	"forum/common"
	"forum/model"
	"forum/service"
	"log"

	"github.com/gin-gonic/gin"
)

//HandleAddCategory 添加分类
func HandleAddCategory(ctx *gin.Context) {
	var category model.Category
	_=ctx.BindJSON(&category)
	log.Println(category)
	ok,err:=service.AddCategoryService(category);
	if !ok{
		common.HandleServerError(ctx,err)
		return
	}
	if ok&&err!=nil{
		common.HandleRecordExists(ctx)
		return
	}
	common.HandleOkReturn(ctx,nil)
}