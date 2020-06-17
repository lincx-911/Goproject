package main

import (
	_ "miniprogram/database"
	orm "miniprogram/database"
	"miniprogram/routes"
)

// @title Golang Esign API
// @version 1.0
// @description  Golang api of demo
// @termsOfService http://github.com

// @contact.name API Support
// @contact.url http://www.cnblogs.com
// @contact.email ×××@qq.com

//@host 127.0.0.1:8080
func main() {
	defer orm.DB.Close()
	route := routes.InitRouter()
	route.Static("/img", "./static/img")
	route.Static("/file", "./static/file")
	route.Run(":8080")
}
