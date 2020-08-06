package main

import (
	post "microproject/proto/post"
	user "microproject/proto/user"
	handler "microproject/service/handle"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)


func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("lastest"),
	)
	service.Init()

	user.RegisterUserHandler(service.Server(), new(handler.User))
	post.RegisterPostHandler(service.Server(), new(handler.Post))

	if err:=service.Run();err!=nil{
		logger.Fatal(err)
	}
}