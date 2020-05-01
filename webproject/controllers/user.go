package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	swt "webproject/middleswt"
	"webproject/models"
)


//UserDel 用户删除
func UserDel(w http.ResponseWriter, r *http.Request) {
	
	// var id struct {
	// 	ID int `json:"id"`
	// }
	// fmt.Println("删除blog")
	// err := json.NewDecoder(r.Body).Decode(&id)
	// fmt.Println("id=", id.ID)
	id := 1;
	var err error
	if err != nil {
		fmt.Println("用户解析json出错")
		res := models.Response{Code: http.StatusNotAcceptable, Msg: "删除失败", Data: nil}
		swt.ResponseWithJson(w, http.StatusNotAcceptable, res)
		return
	}
	err = models.DeleteUser(id)
	if err != nil {
		fmt.Println("删除错误",err)
		res := models.Response{Code: http.StatusServiceUnavailable, Msg: "删除失败", Data: nil}
		swt.ResponseWithJson(w, http.StatusServiceUnavailable, res)
	} else {
		fmt.Println("删除成功",err)
		res := models.Response{Code: http.StatusOK, Msg: "删除成功", Data: nil}
		swt.ResponseWithJson(w, http.StatusOK, res)
	}
}

//UserList 用户列表
func UserList(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/emp/userlist.gtpl", "views/commons/bar.gtpl"))
	users, err := models.GetAllUser()
	if err != nil {
		res := models.Response{Code: http.StatusNotAcceptable, Msg: "用户列表获取失败", Data: nil}
		swt.ResponseWithJson(w, http.StatusNotAcceptable, res)
		users = nil
	}
	t.ExecuteTemplate(w, "userlist", users)
}
