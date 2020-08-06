package main

import (
	"fmt"
	_"github.com/gin-gonic/gin"
	m "miniprogram/models"
)

func main()  {
	var user m.User
	user.UID=1
	fmt.Println(user)
	// r:=gin.Default()
	// r.GET("/ping",func (c *gin.Context)  {
	// 	c.JSON(200,gin.H{
	// 		"message":"pong",
	// 	})
	// })
	// r.Run()
}