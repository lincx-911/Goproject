package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	swt "webproject/middleswt"
	"webproject/models"
)

//Bloglogin 博客登录
func Bloglogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header["Cookie"])
	t, _ := template.ParseFiles("login")
	t.Execute(w, nil)

}

//Blogindex 博客首页
func Blogindex(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("views/blogs/blog.gtpl", "views/blogs/narbar/narbar.gtpl"))
	blogs,err:=models.GetAllBlog()
	if err!=nil{
		res:=models.Response{Code:http.StatusNotAcceptable,Msg: "blog列表获取失败",Data: nil}
		swt.ResponseWithJson(w,http.StatusNotAcceptable,res)
		blogs=nil
	}
	fmt.Println("blog:",blogs)
	t.ExecuteTemplate(w, "blog", blogs)
}

//Blogarticle 文章
func Blogarticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idstr:=r.FormValue("id")
	id,_:=strconv.Atoi(idstr)
	fmt.Println("id",id)
	blog,err:=models.GetBlogbyID(id)
	if err!=nil{
		fmt.Println("失败")
	}
	fmt.Println(blog)
	t := template.Must(template.ParseFiles("views/blogs/article.gtpl", "views/blogs/narbar/narbar.gtpl"))
	t.ExecuteTemplate(w, "article", blog)
}

//ErrorDel 错误处理
func ErrorDel(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("views/error.gtpl")
	t.Execute(w, nil)
}
