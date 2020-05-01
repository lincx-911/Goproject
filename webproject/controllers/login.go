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
	fmt.Println("Register", r.Method)
	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		fmt.Println(user.Username)
		fmt.Println(user.Password)
		if err != nil {
			swt.ResponseWithJson(w, http.StatusBadRequest,
				models.Response{Code: http.StatusBadRequest, Msg: "bad params"})
			return
		}
		x, err := models.GetUserbyName(user.Username)
		if x == user {
			swt.ResponseWithJson(w, http.StatusForbidden,
				models.Response{Code: http.StatusForbidden, Msg: "用户已存在"})
			return
		}
		_, err = models.Insert(user)
		if err != nil {
			swt.ResponseWithJson(w, http.StatusInternalServerError,
				models.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
			return
		}
		swt.ResponseWithJson(w, http.StatusOK,
			models.Response{Code: http.StatusOK, Msg: "注册成功"})
		return
	}
	t, _ := template.ParseFiles("views/regin.gtpl")
	t.Execute(w, nil)
}

//Login 用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/login.gtpl")
		t.Execute(w, nil)
	} else {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			swt.ResponseWithJson(w, http.StatusBadRequest,
				models.Response{Code: http.StatusBadRequest, Msg: "bad params",Data: nil})
		}
		
		fmt.Println(user.Username,user.Password)
		user1, err := models.GetUserbyName(user.Username)
		fmt.Println(user1)
		if err == nil {
			if user.Password == user1.Password {
				token, _ := swt.GenerateToken(&user)
				fmt.Println("token:",token)
				swt.ResponseWithJson(w, http.StatusOK,
					models.Response{Code: http.StatusOK, Msg: "登陆成功",Data: token})

			} else {

				swt.ResponseWithJson(w, http.StatusNotFound,
					models.Response{Code: http.StatusNotFound, Msg: "the password not right"})
			}

		} else {
			swt.ResponseWithJson(w, http.StatusNotFound,
				models.Response{Code: http.StatusNotFound, Msg: "the user not exist"})
		}
	}

}

//SignOut 注销处理
func SignOut(w http.ResponseWriter, r *http.Request) {
	r.Header.Del("Cookie")
	http.Redirect(w,r,"/login",http.StatusFound)
}
