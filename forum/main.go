package main

import (
	"github.com/gin-gonic/gin" 
	orm "forum/respository"
	"forum/router"
)

func main() {
	defer orm.DB.Close()
	gin.SetMode(gin.DebugMode)
	app:=router.IninRouter()
	app.Static("/img", "./static/img")
	app.Run(":8082")
}