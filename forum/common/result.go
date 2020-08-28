package common

import (
	"net/http"
	"github.com/gin-gonic/gin")

//Result 请求响应体 
type Result struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

//HandleAPIReturn 返回请求
func HandleAPIReturn(ctx *gin.Context,code int,msg string,data interface{}) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":data,
	})
}
//HandleOkReturn 请求成功
func HandleOkReturn(ctx *gin.Context,data interface{}){
	HandleAPIReturn(ctx,OK,"Success！",data)
}
//HandleParamsError 传参错误
func HandleParamsError(ctx *gin.Context)  {
	HandleAPIReturn(ctx,ParamsError,"Params Error!",nil)
}
//HandleServerError 服务器内部错误
func HandleServerError(ctx *gin.Context,err interface{}){
	HandleAPIReturn(ctx,ServerError,"Server Error!",err)
}
//HandleOperationError 操作错误
func HandleOperationError(ctx *gin.Context,err interface{}){
	HandleAPIReturn(ctx,OperationError,"Operation Error!",err)
}
//HandleRecordExists 记录已存在
func HandleRecordExists(ctx *gin.Context){
	HandleAPIReturn(ctx,RecordExists,"Record had Existed!",nil)
}
//HandleRecordNotFound 记录找不到
func HandleRecordNotFound(ctx *gin.Context){
	HandleAPIReturn(ctx,RecordNotFound,"Record Not Found!",nil)
}
//HandleAuthError 缺少权限
func HandleAuthError(ctx *gin.Context){
	HandleAPIReturn(ctx,AuthError,"登录超时请重新登录!",nil)
}

	