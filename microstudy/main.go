package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-micro/v2/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
        op.Addrs = []string{"http://127.0.0.1:2379"}
    })
	router := gin.Default()
	router.Handle("GET", "/my", func(c *gin.Context) {
		c.String(http.StatusOK, "my api")
	})
	server := web.NewService(
		web.Name("prodservice"),
		web.Address(":8000"), 
		web.Handler(router),
		web.Registry(consultReg),
	)
	server.Run()
}
