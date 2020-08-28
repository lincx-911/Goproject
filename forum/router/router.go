package router

import (
	"forum/api"
	"forum/middleware"

	"github.com/gin-gonic/gin"
)

//IninRouter 初始化路由
func IninRouter() *gin.Engine {
	router := gin.Default()
	ruser := router.Group("/user")
	{
		ruser.GET("/getself",middleware.Cors(),middleware.AuthMiddleWare(),api.GetUserBySelfHandle)
		ruser.POST("/test",middleware.AuthMiddleWare(),api.AddArticleHandler)
		ruser.GET("/get/:id",middleware.AuthMiddleWare(),api.GetUserByIDHandle)
		ruser.PUT("/update",middleware.AuthMiddleWare(),api.UpdateUserHandle)
		ruser.POST("/updateimg",middleware.AuthMiddleWare(),api.UserUpdateImg)
		ruser.GET("/regist/sendemail", api.Sendemail)
		ruser.PUT("/updatepsw",api.UpdatePsw)
	}
	rcategory:=router.Group("/cate")
	{
		rcategory.POST("/add",api.HandleAddCategory)
	}
	rarticle:=router.Group("/article")
	{
		rarticle.GET("/get/:id",middleware.AuthMiddleWare(),api.GetArticleByIDHandle)
		rarticle.POST("/add",middleware.Cors(),middleware.AuthMiddleWare(),api.AddArticleHandler)
		rarticle.DELETE("/delete/:id",middleware.AuthMiddleWare(),api.DeleteArticleHandle)
		rarticle.PUT("/update",middleware.AuthMiddleWare(),api.UpdateArticleHandle)
		rarticle.GET("/getlist/:index",middleware.AuthMiddleWare(),api.GetArticleListHandle)
		rarticle.GET("/getlists/time/:index",middleware.AuthMiddleWare(),api.GetArticleListByTimeHandle)
		rarticle.GET("/getlists/com/:index",middleware.AuthMiddleWare(),api.GetArticleByComnumHandle)
		rarticle.GET("/getlists/like/:index",middleware.AuthMiddleWare(),api.GetArticleByLikeHandle)
		rarticle.GET("/bycateid/getlist/:cid/:index",middleware.AuthMiddleWare(),api.GetCateArticleListHandle)//通过cateid返回默认文章列表
		rarticle.GET("/bycateid/getlists/time/:cid/:index",middleware.AuthMiddleWare(),api.GetCateArticleListByTimeHandle)//.....返回最进发表
		rarticle.GET("/bycateid/getlists/com/:cid/:index",middleware.AuthMiddleWare(),api.GetCateArticleByComnumHandle)//.....评论最多
		rarticle.GET("/bycateid/getlists/like/:cid/:index",middleware.AuthMiddleWare(),api.GetCateArticleByLikeHandle)//...点赞最多
		rarticle.GET("/getlikes",middleware.AuthMiddleWare(),api.GetUserLikeArticleHandle)
		rarticle.PUT("/likeup/:id",middleware.AuthMiddleWare(),api.LikeNumUpHandle)
		rarticle.PUT("/likedown/:id",middleware.AuthMiddleWare(),api.LikeNumDownHandle)
		rarticle.GET("/getbyuser/:uid/:index",middleware.AuthMiddleWare(),api.GetArticleByUserID)
		rarticle.GET("/getbyself/own/:index",middleware.AuthMiddleWare(),api.GetArticleBySlef)
		rarticle.GET("/getfocus/own/:index",middleware.AuthMiddleWare(),api.GetUserFocueArticleHandle)
		rarticle.GET("/judge/:id",middleware.AuthMiddleWare(),api.JudgeFocusHandle)
		rarticle.PUT("/addfocus/:id",middleware.AuthMiddleWare(),api.AddFocusHandle)
		rarticle.DELETE("/delfocus/:id",middleware.AuthMiddleWare(),api.DeleteFocusHandle)
	}
	rcomment:=router.Group("/comment")
	{
		rcomment.POST("/add",middleware.AuthMiddleWare(),api.AddCommentHandler)
		rcomment.DELETE("/delete/:id",middleware.AuthMiddleWare(),api.DeleteCommentHandle)
		rcomment.PUT("/likeup/:id",middleware.AuthMiddleWare(),api.CommentLikeNumUpHandle)
		rcomment.PUT("/likedown/:id",middleware.AuthMiddleWare(),api.CommentLikeNumDownHandle)
		rcomment.GET("/get/:rid/:index",middleware.AuthMiddleWare(),api.GetCommentListByRaid)
		rcomment.GET("/getlikes",middleware.AuthMiddleWare(),api.GetUserLikeCommentHandle)
		rcomment.GET("/getown/own/:index",middleware.AuthMiddleWare(),api.GetCommentListBySelf)
	}

	return router
}