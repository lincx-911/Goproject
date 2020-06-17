package routes

import (
	"miniprogram/apis"
	//"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	radmin := router.Group("/admin")
	{
		radmin.GET("/getadmins", apis.GetallAdmins)
		radmin.POST("/login", apis.AdminLogin)
		radmin.POST("/addadmin", apis.AdminAdd)
		radmin.DELETE("/deladmin/:id", apis.AdminDel)
		radmin.PUT("/updateadmin/:id", apis.UpdateAdmin)
	}
	ruser := router.Group("/user")
	{
		ruser.POST("/regist", apis.UserRegist)
		ruser.POST("/login", apis.UserLogin)
		ruser.GET("/getallusers", apis.GetallUsers)
		ruser.DELETE("/deleteuser/:id", apis.UserDel)
		ruser.PUT("/updateuser/:id", apis.UpdateUser)
		ruser.GET("/getmyrecruir/:id", apis.GetmyRecruir)
		ruser.GET("/getmyrecruirs", apis.GetallCollectre)        //获取个人发布的招募信息
		ruser.POST("/issuerecruit/:id", apis.IssueRecruit)       //发布招募信息
		ruser.PUT("/aditrecruit/:id", apis.AditRecruit)          //编辑招募信息
		ruser.DELETE("/deleterecruit/:id", apis.DeleteRecruit)   //删除招募信息
		ruser.POST("/addcollectcon", apis.AddCollectcon)         //添加收藏的赛事信息
		ruser.DELETE("/deletecolectcon", apis.DeleteColectcon)   //删除收藏的赛事信息
		ruser.GET("/getallcllectcon/:id", apis.GetallCollectcon) //获取全部收藏赛事收藏
		ruser.GET("/getallcollectre/:id", apis.GetallCollectre)  //获取全部收藏的招募信息
		ruser.POST("/addcollectre", apis.AddCollectre)           //添加收藏的招募信息
		ruser.DELETE("/deletecolectre", apis.DeleteColectre)     //删除收藏的招募信息
		ruser.POST("/uploaduimg/:id", apis.Saveuimg)
		
	}
	rorg := router.Group("/org")
	{
		rorg.POST("/regist", apis.OrganizeRegist)
		rorg.POST("/login", apis.OganizeLogin)
		rorg.DELETE("/delete/:id", apis.OrganizeDel)
		rorg.PUT("/update/:id", apis.OrganizeUpdate)
		rorg.GET("/getorganize/:id", apis.GetOrgbyid)
		rorg.GET("/organizes", apis.GetallOrganizes)
		rorg.POST("/contest/issue", apis.IssueContest)
		rorg.PUT("/contest/update", apis.AdditContest)
		rorg.DELETE("/contest/delete/:id", apis.DeleteContest)
		rorg.GET("/contestbyid/:id", apis.Getcontest)
		rorg.GET("/contestbytag", apis.Getcontestfromtag)
		rorg.GET("/conbypub", apis.GetcontestfromPublic)
		rorg.GET("/contests", apis.Getcontests)
		rorg.POST("/uploadoimg/:id", apis.Saveoimg)
		rorg.POST("/uploadconfile/:id",apis.Uploadconfile)
	}

	//router.StaticFS("/img",http.Dir("../static/img"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
