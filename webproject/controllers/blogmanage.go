package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	swt "webproject/middleswt"
	"webproject/models"
	"webproject/utils"
)

//Manageindex 后台管理首页
func Manageindex(w http.ResponseWriter,r *http.Request){
	t := template.Must(template.ParseFiles("views/manageindex.gtpl", "views/commons/bar.gtpl"))
 	t.ExecuteTemplate(w, "model", "")
}
//Bloglist 博客列表
func Bloglist(w http.ResponseWriter,r *http.Request){
	t := template.Must(template.ParseFiles("views/emp/bloglist.gtpl", "views/commons/bar.gtpl"))
	blogs,err:=models.GetAllBlog()
	if err!=nil{
		res:=models.Response{Code:http.StatusNotAcceptable,Msg: "blog列表获取失败",Data: nil}
		swt.ResponseWithJson(w,http.StatusNotAcceptable,res)
		blogs=nil
	}
	t.ExecuteTemplate(w, "list", blogs)
}
//Blogadd 添加博客
func Blogadd(w http.ResponseWriter,r *http.Request){
	fmt.Println("添加blog",r.Method)
	if r.Method=="POST"{
		// fmt.Println(r.Header)
		var blog models.Blog
		err := json.NewDecoder(r.Body).Decode(&blog)
		fmt.Println(blog,err)
		utils.CheckError(err)
		err=models.InsertBg(blog)
		if err!=nil{
			fmt.Println("数据库错误",err)
			res:=models.Response{Code:http.StatusForbidden,Msg: "添加失败",Data: nil}
			fmt.Println("res:",res)
			swt.ResponseWithJson(w,http.StatusForbidden,res)
			return
		}
		fmt.Println("数据库正确")
		res:=models.Response{Code:http.StatusOK,Msg: "添加成功",Data: nil}//返回给前端的响应体
		swt.ResponseWithJson(w,http.StatusOK,res)//将res的内容写入w
	}else{
		t := template.Must(template.ParseFiles("views/emp/addadit.gtpl", "views/commons/bar.gtpl"))
		t.ExecuteTemplate(w, "addadit", nil)
	}
}
//Blogdel 删除博客
func Blogdel(w http.ResponseWriter,r *http.Request){
	var id struct{
		ID string `json:"id"`
	}
	fmt.Println("删除blog")
	err := json.NewDecoder(r.Body).Decode(&id)
	if err!=nil{
		fmt.Println("blog添加解析json出错")
		res:=models.Response{Code:http.StatusNotAcceptable,Msg: "添加失败",Data: nil}
		swt.ResponseWithJson(w,http.StatusNotAcceptable,res)
		return
	}
	fmt.Println("id=",id.ID)
	err=models.DeleteBg(id.ID)
	if err!=nil{
		res:=models.Response{Code:http.StatusServiceUnavailable,Msg: "删除失败",Data:nil}
		swt.ResponseWithJson(w,http.StatusServiceUnavailable,res)
	}else{
		res:=models.Response{Code:http.StatusOK,Msg: "删除成功",Data: nil}
		swt.ResponseWithJson(w,http.StatusOK,res)
	}
}
//Blogadit 编辑博客
func Blogadit(w http.ResponseWriter,r *http.Request){
	fmt.Println("编辑blog",r.Method)
	if r.Method=="POST"{
		fmt.Println(r.Header)
		var blog models.Blog
		err := json.NewDecoder(r.Body).Decode(&blog)
		fmt.Println(blog,err)
		err=models.UpdateBg(blog)
		if err!=nil{
			res:=models.Response{Code:http.StatusServiceUnavailable,Msg: "修改成功",Data: nil}
			swt.ResponseWithJson(w,http.StatusServiceUnavailable,res)
			return
		}
		res:=models.Response{Code:http.StatusOK,Msg: "修改成功",Data: nil}
		swt.ResponseWithJson(w,http.StatusOK,res)
	}else{
		r.ParseForm()
		id1:=r.FormValue("id")
		id,err:=strconv.Atoi(id1)
		fmt.Println(id)
		b,err:=models.GetBlogbyID(id)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(b)
		t := template.Must(template.ParseFiles("views/emp/addadit.gtpl", "views/commons/bar.gtpl"))
		t.ExecuteTemplate(w, "addadit", b)
	}
}
