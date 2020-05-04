package main

import (
	"fmt"
	"net/http"
	cs "webproject/controllers"
	swt "webproject/middleswt"
)
// func test(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w,"hello world")
// }
func main() {
	fmt.Println("访问服务器了")
	//静态资源的引用  我之前是用的 /statics/ 然后模板中写的是 /statics/css/然后访问不到css
	//所以路由应该需要指到具体文件的当前文件夹名   eg: /statics/css/1.css  如果渲染1.css 需注册路由/statics/css/
	http.Handle("/statics/css/", http.StripPrefix("/statics/css/", http.FileServer(http.Dir("statics/css/"))))
	http.Handle("/statics/js/", http.StripPrefix("/statics/js/", http.FileServer(http.Dir("statics/js/"))))
	http.Handle("/statics/image/", http.StripPrefix("/statics/image/", http.FileServer(http.Dir("statics/image/"))))
	http.HandleFunc("/reginster", cs.Register)
	http.HandleFunc("/login", cs.Login)
	http.HandleFunc("/blogindex", cs.Blogindex)
	http.HandleFunc("/blogarticle", cs.Blogarticle)
	http.HandleFunc("/bloglist", swt.TokenMiddleware(cs.Bloglist))
	http.HandleFunc("/index", cs.Manageindex)
	http.HandleFunc("/blogadd", swt.TokenMiddleware(cs.Blogadd))
	http.HandleFunc("/blogdelete", cs.Blogdel)
	http.HandleFunc("/blogadit", swt.TokenMiddleware(cs.Blogadit))
	http.HandleFunc("/error", cs.ErrorDel)
	http.HandleFunc("/userlist",swt.TokenMiddleware(cs.UserList))
	http.HandleFunc("/userdelete",cs.UserDel)
	http.HandleFunc("/signout",cs.SignOut)
	// http.HandleFunc("/test",test)
	http.ListenAndServe(":8080", nil)
}
