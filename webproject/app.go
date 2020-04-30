package main

import (
	"fmt"
	"net/http"
	cs "webproject/controllers"
	swt "webproject/middleswt"
)

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
	http.HandleFunc("/blogarticle", swt.TokenMiddleware(cs.Blogarticle))
	http.HandleFunc("/bloglist", cs.Bloglist)
	http.HandleFunc("/index", cs.Manageindex)
	http.HandleFunc("/blogadd", cs.Blogadd)
	http.HandleFunc("/blogdelete",cs.Blogdel)
	http.HandleFunc("/blogadit",cs.Blogadit)
	http.HandleFunc("/test",cs.Test)
	http.ListenAndServe(":8080", nil)
}
