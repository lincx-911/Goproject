package api

import (
	"fmt"
	"forum/common"
	"forum/model"
	"forum/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//AddCommentHandler 发表评论
func AddCommentHandler(ctx *gin.Context) {
	uid := ctx.GetInt("uid")
	fmt.Println("uid", uid)
	var comment model.Comment
	_ = ctx.ShouldBindJSON(&comment)
	comment.UserID=uid
	fmt.Println("comment",comment)
	if comment.Content == "" {
		common.HandleParamsError(ctx)
		return
	}
	_, err := service.AddCommentService(&comment)
	if err != nil {
		common.HandleOperationError(ctx, err)
		return
	}
	println("commentid",comment.ID )
	common.HandleOkReturn(ctx, comment)
}

//DeleteCommentHandle 通过id删除评论
func DeleteCommentHandle(ctx *gin.Context)  {
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	fmt.Println("aid",aid)
	if ok,err:=service.DeleteCommentByID(aid);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}
//CommentLikeNumUpHandle 文章点赞数增加
func CommentLikeNumUpHandle(ctx *gin.Context)  {
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	uid:=ctx.GetInt("uid")
	if uid==0||aid==0{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.ComLikenumUpService(aid,uid);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}
//CommentLikeNumDownHandle 文章点赞数减少
func CommentLikeNumDownHandle(ctx *gin.Context)  {
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	uid:=ctx.GetInt("uid")
	if uid==0||aid==0{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.ComLikenumDownService(aid,uid);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}
//GetCommentListByRaid 获取文章的全部评论
func GetCommentListByRaid(ctx *gin.Context){
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	rid,err:=strconv.Atoi(ctx.Params.ByName("rid"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	commentlist,totalnum,err:=service.GetCommentByRid(rid,index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["comments"]=commentlist
	common.HandleOkReturn(ctx,result)
}
//GetUserLikeCommentHandle 获取用户点赞的评论id
func GetUserLikeCommentHandle(ctx *gin.Context)  {
	uid:=ctx.GetInt("uid")
	if uid==0{
		common.HandleAuthError(ctx)
		return
	}
	alist,err:=service.GetUserLikeCommentService(uid)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,alist)
}
//GetCommentListBySelf 获取本人的全部评论
func GetCommentListBySelf(ctx *gin.Context){
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	uid:=ctx.GetInt("uid")
	if uid==0{
		common.HandleParamsError(ctx)
		return
	}
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	commentlist,totalnum,err:=service.GetUserCommentByLimitService(uid,index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["comments"]=commentlist
	common.HandleOkReturn(ctx,result)
}