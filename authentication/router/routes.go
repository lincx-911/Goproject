package router

import (
	"authentication/controller"
	"authentication/middleware"

	"github.com/gorilla/mux"
)

//InitRouter 初始化路由
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	rjwt := router.PathPrefix("/api").Subrouter()
	rjwt.HandleFunc("/test", controller.TestFunc).Methods("GET")
	rjwt.Use(middleware.AuthMiddleware)
	router.HandleFunc("/gettoken", controller.GetToken).Methods("GET")
	return router
}
