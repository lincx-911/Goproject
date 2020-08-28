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

//AddArticleHandler 发布文章
func AddArticleHandler(ctx *gin.Context) {
	var article model.Article
	_=ctx.ShouldBind(&article)
	if article.Title==""||article.Content==""{
		common.HandleParamsError(ctx)
		return
	}
	uid := ctx.GetInt("uid")
	fmt.Println("uid", uid)
	article.UserID=uid
	id, err := service.AddArticleService(&article)
	if err != nil {
		common.HandleOperationError(ctx, err)
		return
	}
	article.ID = id
	println("articleid", id)
	
	common.HandleOkReturn(ctx, article)
	return
}

//GetArticleByIDHandle 通过id获取文章
func GetArticleByIDHandle(ctx *gin.Context){
	
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	article,err:=service.GetArticleByIDServie(aid)
	if err!=nil{
		common.HandleOperationError(ctx,err)
		return
	}
	if article.ID==0{
		common.HandleRecordNotFound(ctx)
		return
	}

	common.HandleOkReturn(ctx,article)
	return 
}
//DeleteArticleHandle 通过id删除文章
func DeleteArticleHandle(ctx *gin.Context)  {
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	fmt.Println("aid",aid)
	if ok,err:=service.DeleteArticleByID(aid);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}

//UpdateArticleHandle 更新文章
func UpdateArticleHandle(ctx *gin.Context) {
	var article model.Article
	_ = ctx.BindJSON(&article)
	if article.Title==""||article.Content==""{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.UpdateArticleService(&article);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}

//LikeNumUpHandle 文章点赞数增加
func LikeNumUpHandle(ctx *gin.Context)  {
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	uid := ctx.GetInt("uid")
	fmt.Println("aid",aid)
	if aid==0||uid==0{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.LikenumUpService(aid,uid);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}
//LikeNumDownHandle 文章点赞数减少
func LikeNumDownHandle(ctx *gin.Context)  {
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	uid := ctx.GetInt("uid")
	if aid==0||uid==0{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.LikenumDownService(aid,uid);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}
//GetArticleListHandle 获取文章列表默认方式
func GetArticleListHandle(ctx *gin.Context)  {
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	fmt.Println(index)
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}

	articlelist,totalnum,err:=service.GetArticleListService(index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}

//GetArticleListByTimeHandle 获取文章列表时间排序方式
func GetArticleListByTimeHandle(ctx *gin.Context)  {
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}

	articlelist,totalnum,err:=service.GetArticleByTimeService(index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}
//GetArticleByComnumHandle 获取评论数文章列表
func GetArticleByComnumHandle(ctx *gin.Context)  {
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}

	articlelist,totalnum,err:=service.GetArticleByComService(index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}
//GetArticleByLikeHandle 获取文章列表按点赞数
func GetArticleByLikeHandle(ctx *gin.Context)  {
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}

	articlelist,totalnum,err:=service.GetArticleByLikeService(index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}
//AddFocusHandle 添加文章关注
func AddFocusHandle(ctx *gin.Context)  {
	uid := ctx.GetInt("uid")
	fmt.Println("uid", uid)
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	if aid==0||uid==0||err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.AddFocusArticleService(uid,aid);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}
//DeleteFocusHandle 取消文章关注
func DeleteFocusHandle(ctx *gin.Context)  {
	uid := ctx.GetInt("uid")
	fmt.Println("uid", uid)
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.DeleteFocusArticleService(uid,aid);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}
//JudgeFocusHandle 判断是否关注
func JudgeFocusHandle(ctx *gin.Context)  {
	uid := ctx.GetInt("uid")
	fmt.Println("uid", uid)
	aid,err:=strconv.Atoi(ctx.Params.ByName("id"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.JudgeFocusService(uid,aid);!ok{
		if err==nil{
			common.HandleOkReturn(ctx,0)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,1)
}
//GetArticleBySlef 获取个人的文章
func GetArticleBySlef(ctx *gin.Context)  {
	uid := ctx.GetInt("uid")
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	articlelist,totalnum,err:=service.GetArticleByUserIDService(uid,index,common.PAGE_SIZE)
	
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}
//GetArticleByUserID 获取user发表的全部文章
func GetArticleByUserID(ctx *gin.Context)  {
	uid,err:=strconv.Atoi(ctx.Params.ByName("uid"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	articlelist,totalnum,err:=service.GetArticleByUserIDService(uid,index,common.PAGE_SIZE)
	
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}
//GetUserLikeArticleHandle 获取用户点赞的文章id
func GetUserLikeArticleHandle(ctx *gin.Context)  {
	uid:=ctx.GetInt("uid")
	if uid==0{
		common.HandleAuthError(ctx)
		return
	}
	alist,err:=service.GetUserLikeArticleService(uid)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,alist)
}
//GetCateArticleListHandle 获取文章列表默认方式
func GetCateArticleListHandle(ctx *gin.Context)  {
	cid,err:=strconv.Atoi(ctx.Params.ByName("cid"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	fmt.Println(index)
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}

	articlelist,totalnum,err:=service.GetCateArticleListService(cid,index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}

//GetCateArticleListByTimeHandle 获取文章列表时间排序方式
func GetCateArticleListByTimeHandle(ctx *gin.Context)  {
	cid,err:=strconv.Atoi(ctx.Params.ByName("cid"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}

	articlelist,totalnum,err:=service.GetCateArticleByTimeService(cid,index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}
//GetCateArticleByComnumHandle 获取评论数文章列表
func GetCateArticleByComnumHandle(ctx *gin.Context)  {
	cid,err:=strconv.Atoi(ctx.Params.ByName("cid"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}

	articlelist,totalnum,err:=service.GetCateArticleByComService(cid,index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}
//GetCateArticleByLikeHandle 获取文章列表按点赞数
func GetCateArticleByLikeHandle(ctx *gin.Context)  {
	cid,err:=strconv.Atoi(ctx.Params.ByName("cid"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}

	articlelist,totalnum,err:=service.GetCateArticleByLikeService(cid,index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}
//GetUserFocueArticleHandle 获取本人关注的文章列表
func GetUserFocueArticleHandle(ctx *gin.Context)  {
	uid := ctx.GetInt("uid")
	index,err:=strconv.Atoi(ctx.Params.ByName("index"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	articlelist,totalnum,err:=service.GetUserFocueArticleService(uid,index,common.PAGE_SIZE)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	result:=make(map[string]interface{})
	result["totalnum"]=totalnum
	result["articles"]=articlelist
	common.HandleOkReturn(ctx,result)
}