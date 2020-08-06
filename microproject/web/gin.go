// package main

// import (
// 	"context"
// 	user "microproject/proto/user"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	micro "github.com/micro/go-micro/v2"
// 	"github.com/micro/go-micro/v2/logger"
// )

// type Say struct{}

// var (
// 	cl user.UserService
// )

// func (s *Say)Anything(c *gin.Context)  {
// 	logger.Info("Received Say.Anything API request")
// 	c.JSON(http.StatusOK,map[string]string{
// 		"message":"This is the microproject API",
// 	})
// }

// func (s *Say)Hello(c *gin.Context){
// 	logger.Info("Received Say.Hello API request")
// 	name:=c.Param("name")
	
// 	response,err:=cl.QueryUserByName(context.TODO(),&user.Request{
// 		UserName:name,
// 	})
// 	if err!=nil{
// 		logger.Error(err)
// 		c.JSON(http.StatusInternalServerError,err)
// 	}
// 	c.JSON(http.StatusOK,response)
// }

// func main()  {
// 	service:=micro.NewService(micro.Name("web.gin"))
// 	service.Init()
// 	cl = user.NewUserService("go.micro.srv.user",service.Client())

// 	say := new(Say)
// 	router:=gin.Default()
// 	router.GET("/greeter",say.Anything)
// 	router.GET("/greeter/:name",say.Hello)
// 	router.Run(":8081")
// }