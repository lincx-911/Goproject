package controller

import (
	"authentication/common"
	"encoding/json"
	"fmt"
	"net/http"
)

//GetToken 获取token
func GetToken(w http.ResponseWriter,r *http.Request)  {
	// var appclaim struct{
	// 	Appname string `json:"appname"`
	// 	Appinfo string `json:"appinfo"`
	// }
	//appclaim:=make(map[string]interface{})
	var appclaim common.MyClaim
	if err:=json.NewDecoder(r.Body).Decode(&appclaim);err!=nil{
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w,"服务器错误")
		return
	}
	if token,err:=common.GenToken(appclaim);err!=nil{
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w,"服务器错误")
		
	}else{
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,token)
	}
	
}

//TestFunc 测试
func TestFunc(w http.ResponseWriter,r *http.Request)  {
	var approles []int
	if err:=json.NewDecoder(r.Body).Decode(&approles);err!=nil{
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w,"服务器错误")
		return
	}
	fmt.Println(len(approles))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w,approles)
}