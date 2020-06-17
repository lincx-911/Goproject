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
		ruser.POST("/regist", apis.UserRegist)             //注册
		ruser.POST("/login", apis.UserLogin)               //登录
		ruser.GET("/getallusers", apis.GetallUsers)        //获取用户列表
		ruser.DELETE("/deleteuser/:id", apis.UserDel)      //删除用户
		ruser.PUT("/updateuser", apis.UpdateUser)          //更新用户信息
		ruser.GET("/getuser/:id", apis.SelectUserfid)      //用过学号查找用户
		ruser.GET("/myinssue/:id", apis.Getperinssue)      //获取个人发布的信息
		ruser.PUT("/comfiretask", apis.Comfiretask)        //确认任务
		ruser.PUT("/delcomfire", apis.Delsuretask)         //取消确认
		ruser.PUT("/changeimg/:id", apis.Saveuimg)         //更新头像
		ruser.POST("/regist/sendemail", apis.Sendemail)    //发送邮箱验证
		ruser.PUT("/update/password", apis.Updateuserpass) //修改密码
	}
	rinfo := router.Group("/info")
	{
		rinfo.POST("/inssue", apis.IssueInfo)              //发布招募信息
		rinfo.PUT("/aditinfo", apis.AdditInfo)             //编辑招募信息
		rinfo.DELETE("/delinfo/:id", apis.DeleteInfo)      //删除招募信息
		rinfo.GET("/getinfo/:id", apis.GetinfobyID)        //通过id查找信息
		rinfo.GET("/getinssuetag", apis.Getcontestfromtag) //通过标签获取信息列表
		rinfo.GET("/getinssues", apis.Getcontests)         //获取全部招募信息
		rinfo.GET("/gettaginssues",apis.Getfreecontests)   //按类别获取招募信息
	}

	//router.StaticFS("/img",http.Dir("../static/img"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
