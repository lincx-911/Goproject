package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	swt "webproject/middleswt"
	"webproject/models"
)

//Register 用户注册
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	if err != nil || user.Username == "" || user.Password == "" {
		swt.ResponseWithJson(w, http.StatusBadRequest,
			models.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		return
	}
	_, err = models.Insert(user)
	if err != nil {
		swt.ResponseWithJson(w, http.StatusInternalServerError,
			models.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
	}
}

//Login 用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/login.gtpl")
		t.Execute(w, nil)
	}else{
		r.ParseForm()
		// err := json.NewDecoder(r.Body).Decode(&user)
		// if err != nil {
		// 	swt.ResponseWithJson(w, http.StatusBadRequest,
		// 		models.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		// }
		username:=r.FormValue("username")
		password:=r.FormValue("password")
		user, err := models.GetUserbyName(username)
		if err == nil {
			if user.Password==password{
				token, _ := swt.GenerateToken(&user)
				r.Header.Add("Token",token)
				swt.ResponseWithJson(w, http.StatusOK,
					models.Response{Code: http.StatusOK, Data: models.JwtToken{Token: token}})
				
			}else{
				
				swt.ResponseWithJson(w, http.StatusNotFound,
					models.Response{Code: http.StatusNotFound, Msg: "the password not right"})
			}
			
		} else {
			swt.ResponseWithJson(w, http.StatusNotFound,
				models.Response{Code: http.StatusNotFound, Msg: "the user not exist"})
		}
	}
	
}

//Test 测试
func Test(w http.ResponseWriter,r *http.Request){
	if r.Method=="POST"{
		res:=models.Response{Code:http.StatusOK,Msg: "删除成功",Data: nil}
		swt.ResponseWithJson(w,http.StatusOK,res)
		return
	}
	t, _ := template.ParseFiles("views/test.gtpl")
	B:="hishi"
	t.Execute(w, B)
	
}

