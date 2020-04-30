package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

//Bloglogin 博客登录
func Bloglogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	t, _ := template.ParseFiles("login")
	t.Execute(w, nil)

}

//Blogindex 博客首页
func Blogindex(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	t := template.Must(template.ParseFiles("views/blogs/blog.gtpl", "views/blogs/narbar/narbar.gtpl"))
	t.ExecuteTemplate(w, "blog", nil)
}

//Blogarticle 文章
func Blogarticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	t := template.Must(template.ParseFiles("views/blogs/article.gtpl", "views/blogs/narbar/narbar.gtpl"))
	t.ExecuteTemplate(w, "article", nil)
}
