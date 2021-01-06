package main

import (
	common1 "authentication/common"
	"context"
	"encoding/json"

	//"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

func main() {
	// create a Dapr service (e.g. ":8080", "0.0.0.0:8080", "10.1.1.1:8080" )
	s := daprd.NewService(":8080")

	// add a service to service invocation handler
	if err := s.AddServiceInvocationHandler("/gettoken", gettokenHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
	if err := s.AddServiceInvocationHandler("/test", testHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func gettokenHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf(
		"echo - ContentType:%s, Verb:%s, QueryString:%s, %s",
		in.ContentType, in.Verb, in.QueryString, in.Data,
	)
	var appclaim common1.MyClaim
	//token:=""
	sb := string(in.Data)
	log.Println(sb)
	err = json.Unmarshal(in.Data, &appclaim)
	if err != nil {
		return
	}
	var token1 string
	token1, err = common1.GenToken(appclaim)
	if err != nil {
		return
	}
	res := common1.Result{
		Code: common1.OK,
		Msg:  "ok",
		Data: token1,
	}
	resbuff,_:=json.Marshal(res)
	out = &common.Content{
		Data:       resbuff ,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}

func testHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error){
	if in==nil{
		err = errors.New("invocation parameter required")
		return
	}
	res := common1.Result{
		Code: common1.OK,
		Msg:  "success",
		Data: "林培创最帅",
	}
	resbuff,_:=json.Marshal(res)
	out = &common.Content{
		Data:       resbuff ,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return 
}