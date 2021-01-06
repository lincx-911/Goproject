package controller

import (
	common1 "authentication/common"
	"authentication/model"
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/dapr/go-sdk/service/common"
)


//GettokenHandler 获取token
func GettokenHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf(
		"echo - ContentType:%s, Verb:%s, QueryString:%s, %s",
		in.ContentType, in.Verb, in.QueryString, in.Data,
	)
	var appclaim common1.MyClaim
	sb := string(in.Data)
	log.Println(sb)
	err = json.Unmarshal(in.Data, &appclaim)
	if err != nil {
		return
	}
	var token1 string
	if appclaim.Appname==""||appclaim.Appinfo==""{
		res := common1.Result{
			Code: common1.ParamsError,
			Msg:  "请求参数错误",
			Data: "",
		}
		resbuff,_:=json.Marshal(res)
		out = &common.Content{
			Data:       resbuff ,
			ContentType: in.ContentType,
			DataTypeURL: in.DataTypeURL,
		}
		return
	}
	token1, err = common1.GenToken(appclaim)
	if err != nil {

		return
	}
	result:=make(map[string]interface{})
	result["token"]=token1
	res := common1.Result{
		Code: common1.OK,
		Msg:  "请求成功",
		Data: result,
	}
	resbuff,_:=json.Marshal(res)
	out = &common.Content{
		Data:       resbuff ,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}


//GetStudentsHandler 获取学生信息
func GetStudentsHandler(ctx context.Context,in *common.BindingEvent) (out []byte, err error) {
	token:=in.Metadata["Token"]
	claim, err := common1.ParseToken(token)
	if err!=nil{
		log.Println("学生获取函数发生错误")
		out,err = common1.ResultJSON(400,"登录超时","")
		if err!=nil{
			log.Println("内部错误")
		}
		return out,err
	}
	hasrole:=false 
	for _,value := range claim.Roles{
		
		if value==1{
	
			hasrole=true
			break
		}
	}
	if(!hasrole){
		out,err = common1.ResultJSON(2,"没有该权限","")
		return
	}

	list := []model.Student{
			{
				ID:1,
				Name:"潘晓春",
				Score:98.5,
			},
			{
				ID:2,
				Name:"林小红",
				Score:99.5,
			},
			{
				ID:3,
				Name:"陈小明",
				Score:89.5,
			},
	}

	// res:=make(map[string]interface{})
	// res["code"]=1
	// res["msg"]="请求成功"
	// res["data"]=list
	//out,err=json.Marshal(res)
	out,err = common1.ResultJSON(0,"请求成功",list)
	return 
}
//GetCoursesHandler 获取课程信息
func GetCoursesHandler(ctx context.Context,in *common.BindingEvent) (out []byte, err error) {
	token:=in.Metadata["Token"]
	claim, err := common1.ParseToken(token)
	if err!=nil{
		log.Println("课程获取函数发生错误")
		out,err = common1.ResultJSON(400,"登录超时","")
		if err!=nil{
			log.Println("内部错误")
		}
		return out,err
	}
	hasrole:=false 
	for _,value := range claim.Roles{
		
		if value==1{
	
			hasrole=true
			break
		}
	}
	if(!hasrole){
		out,err = common1.ResultJSON(2,"没有该权限","")
		return
	}

	list := []model.Course{
			{
				ID:1,
				Name:"高等数学",
				Teacher:"黄老师",
			},
			{
				ID:2,
				Name:"通信原理",
				Teacher:"单老师",
			},
			{
				ID:3,
				Name:"计算机网络",
				Teacher:"张老师",
			},
	}
	out,err = common1.ResultJSON(0,"请求成功",list)
	return 
}

//PostPublic 发布公告
func PostPublic(ctx context.Context,in *common.BindingEvent) (out []byte, err error) {
	token:=in.Metadata["Token"]
	claim, err := common1.ParseToken(token)
	if err!=nil{
		log.Println("发布公告函数发生错误")
		out,err = common1.ResultJSON(400,"登录超时","")
		if err!=nil{
			log.Println("内部错误")
		}
		return out,err
	}
	hasrole:=false 
	for _,value := range claim.Roles{
		
		if value==4{
	
			hasrole=true
			break
		}
	}
	if(!hasrole){
		out,err = common1.ResultJSON(2,"没有该权限","")
		return
	}
	out,err = common1.ResultJSON(0,"请求成功",nil)
	return 
}
//GetTeachersHandler 获取教师信息
func GetTeachersHandler(ctx context.Context,in *common.BindingEvent) (out []byte, err error) {
	token:=in.Metadata["Token"]
	claim, err := common1.ParseToken(token)
	if err!=nil{
		log.Println("获取教师信息发生错误")
		out,err = common1.ResultJSON(400,"登录超时","")
		if err!=nil{
			log.Println("内部错误")
		}
		return out,err
	}
	hasrole:=false 
	for _,value := range claim.Roles{
		
		if value==3{
	
			hasrole=true
			break
		}
	}
	if(!hasrole){
		out,err = common1.ResultJSON(2,"没有该权限","")
		return
	}

	list := []model.Teacher{
			{
				ID:1,
				Name:"黄老师",
				Department:"数学系",
			},
			{
				ID:2,
				Name:"单老师",
				Department:"软件系",
			},
			{
				ID:3,
				Name:"张老师",
				Department:"网络系",
			},
	}
	out,err = common1.ResultJSON(0,"请求成功",list)
	return 
}

//RegisterHandler 获取课程信息
func RegisterHandler(ctx context.Context,in *common.InvocationEvent) (out *common.Content, err error) {
	
	sb := string(in.Data)
	log.Println(sb)

	res := common1.Result{
		Code: common1.OK,
		Msg:  "请求成功",
		Data: nil,
	}
	resbuff,_:=json.Marshal(res)
	out = &common.Content{
		Data:       resbuff ,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	
	return
}
//RegisterDevHandler 注册开发者
func RegisterDevHandler(ctx context.Context,in *common.InvocationEvent) (out *common.Content, err error) {
	
	sb := string(in.Data)
	log.Println(sb)
	
	res := common1.Result{
		Code: common1.OK,
		Msg:  "请求成功",
		Data: nil,
	}
	resbuff,_:=json.Marshal(res)
	out = &common.Content{
		Data:       resbuff ,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	
	return
}