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
	Blogs,err:=models.GetAllBlog()

	Tags,err:=models.GetTags()

	Tagges:=make(map[string][]int)
	
	for _,k:= range Tags{
		Tagges[k.Tag]= append(Tagges[k.Tag],k.ID)
	}
	
	Categories,err:=models.GetCategories()

	Cates := make(map[string][]int)
	for _,k:= range Categories{
		Cates[k.Categorie]=append(Cates[k.Categorie],k.ID)
	}
	if err!=nil{
		res:=models.Response{Code:http.StatusNotAcceptable,Msg: "blog列表获取失败",Data: nil}
		swt.ResponseWithJson(w,http.StatusNotAcceptable,res)
		Blogs=nil
	}
	Respose := map[string]interface{}{
		"Blogs":Blogs,
		"Tags":Tagges,
		"Categories":Cates,
	}
	t.ExecuteTemplate(w, "blog",Respose)
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
