package router

import (
	"authentication/controller"
	"log"
	//"authentication/middleware"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	
)
//Server dapr服务
var Server common.Service 
func init(){
	// create a Dapr service (e.g. ":8080", "0.0.0.0:8080", "10.1.1.1:8080" )
	Server = daprd.NewService(":8080")
	if err := Server.AddServiceInvocationHandler("/gettoken", controller.GettokenHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
	if err := Server.AddServiceInvocationHandler("/registapp", controller.RegisterHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
	if err := Server.AddBindingInvocationHandler("/getstudents", controller.GetStudentsHandler); err != nil {
		log.Fatalf("error adding binding handler: %v", err)
	}
	if err := Server.AddBindingInvocationHandler("/getcourses", controller.GetCoursesHandler); err != nil {
		log.Fatalf("error adding binding handler: %v", err)
	}
	if err := Server.AddBindingInvocationHandler("/postpublic", controller.PostPublic); err != nil {
		log.Fatalf("error adding binding handler: %v", err)
	}
	if err := Server.AddBindingInvocationHandler("/getteachers", controller.GetTeachersHandler); err != nil {
		log.Fatalf("error adding binding handler: %v", err)
	}
}