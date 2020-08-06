package main

import (
	"context"
	postProto "microproject/proto/post"
	userProto "microproject/proto/user"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)


func main()  {
	service:=micro.NewService(micro.Name("user.client"))
	service.Init()

	user:=userProto.NewUserService("go.micro.srv.user",service.Client())
	post:=postProto.NewPostService("go.micro.srv.user", service.Client())

	rsp,err:=user.QueryUserByName(context.TODO(),&userProto.Request{UserName: "Tom"})
	if err!=nil{
		logger.Fatal(err)
	}
	logger.Info(rsp.GetUser())

	rsp2, err2 := post.QueryUserPosts(context.TODO(), &postProto.Request{UserId: 1})
	if err2 != nil {
		logger.Fatal(err2)
	}
	// Print response
	logger.Info(rsp2.GetPost().Title)
}