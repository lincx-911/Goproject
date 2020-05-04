
{{define "blog"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>index</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.19.2/axios.js"></script>
    		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
    <link rel="stylesheet" href="statics/css/main.css">
    <script type="text/javascript">
        // function send(){
        //     let token = window.localStorage.getItem("Authorization");
        //     console.log(token);
        //     var currentCookie=document.cookie.split(";")[0];
        //     axios.defaults.headers.common["Authorization"] = token;
        //     console.log(currentCookie);
        //     axios({
		// 		headers: {
		// 			"Authorization":token
		// 		},
        //         methods:'get',
        //         url:'/blogindex',
        //     })
        //         .then(function (res){
		// 			console.log("成功登录");
		// 			axios.defaults.headers.common["Authorization"] = token;
					
		// 		})
		// 		.catch(function (err){
		// 			console.log("post err",err);
		// 			alert(err);
		// 			window.location.reload();
		// 		})
           
            // axios.
            // $.ajax({
            // //几个参数需要注意一下
            //     type: "GET",//方法类型
            //     dataType: "json",//预期服务器返回的数据类型
            //     url: "/blogindex" ,//url
            //     headers:{
            //         'Content-Type':'application/json',
            //         'Token':currentCookie
            //     },
            //     // beforeSend:function(request){
            //     //     request.setRequestHeader("Token",localStorage.token);
            //     // }
            // });
        //}
        //window.onload = send;
    </script>
</head>
<body>
<div class="navbar navbar-default">
    <div class="container">
        <div class="navbar-header">
            <a class="navbar-brand" href="index.html">My Blog</a>
        </div>
        <label class="toggle-label visible-xs-inline-block" for="toggle-checkbox">MENU</label>
        <input class="hidden" id="toggle-checkbox" type="checkbox">
        <div class="hidden-xs">
            <nav class="navbar navbar-default" role="navigation">
                <div class="collapse navbar-collapse">
                <ul class="nav navbar-nav">
                    <li><a href="#">全部分类</a></li>
                    <li><a href="#">语言</a></li>
                    <li><a href="#">数据结构</a></li>
                    <li><a href="#">网络</a></li>
                    <li><a href="#">操作系统</a></li>
                    <li><a href="#">算法</a></li>
                    <li><a href="#">实战</a></li>
                </ul>
            </div>
        </nav>
            <!-- <ul class="nav navbar-nav navbar-right">
                <li><a href="/login">登陆</a></li>
                <li><a href="/reginster">注册</a></li>
            </ul> -->
        </div>
    </div>
</div>

<div class="container">
    <div class="row">
        <div class="col-sm-2 hidden-xs">
            <div class="list-group side-bar">
            
                <a class="list-group-item active" role="button">随笔</a>
                <a class="list-group-item" role="button">随便</a>
                <a class="list-group-item" role="button">随笔</a>
                <a class="list-group-item" role="button">随笔</a>
                <a class="list-group-item" role="button">随笔</a>
                <a class="list-group-item" role="button">随笔</a>
            </div>
        </div>

        <div class="col-sm-7">
            <div class="blog-list">
                <div class="blog-list-item clearfix">
                    <div class="col-xs-5">
                        <img src="statics/image/blog.png">
                    </div>
                    <div class="col-xs-7">
                        {{range .}}
                        <a href="/blogarticle?id={{.ID}}" class="title">
                            {{.Title}}
                        </a>
                        <div class="info">
                            <span class="avatar"><img src="statics/image/v2-2d45613b0fb8cdd36f53c3b31d0c6ee8_hd.jpg"
                                                      alt="avatar"></span>
                            <span>{{.Tag}}</span> |
                            <span>{{.Categorie}}</span>
                            <span class="glyphicon glyphicon-thumbs-up blog-hot" aria-hidden="true"></span> |
                            <span>{{.Date}}</span>
                        </div>
                        {{end}}
                        {{if .}}
                        {{else}}
                        <h2>这个人很懒，啥也没写</h2>
                        {{end}}
                    </div>
                </div>
            </div>
        </div>

        <div class="col-sm-3 hidden-xs">
            <div class="search-bar">
                <form role="form">
                    <div class="form-group has-feedback">
                        <label class="sr-only" for="Search">Search：</label>
                        <input type="search" class="form-control" placeholder="搜索" id="Search">
                        <span class="glyphicon glyphicon-search form-control-feedback"></span>
                    </div>
                </form>
            </div>
            <div class="side-bar-card clearfix">
                <div class="col-xs-5">
                    <img src="statics/image/adver.png">
                </div>
                <div class="col-xs-7 side-bar-introduction">
                    <div class="">代码改变世界</div>
                    <div class="side-bar-phone">联系电话：XXXX</div>
                </div>
            </div>
            <div class="side-bar-card side-bar-recommend clearfix">
                <div class="side-bar-title">推荐阅读</div>
                <div class="side-bar-body">
                    <div class="side-bar-list">
                        <div class="side-bar-item">
                            <a href="blog.html" class="side-item-title">浅析Django项目优化</a>
                            <div class="side-item-info">10.4k阅读 | 五天前</div>
                        </div>
                        <div class="side-bar-item">
                            <a href="blog.html" class="side-item-title">python解释器</a>
                            <div class="side-item-info">0.4k阅读 | 一小时前</div>
                        </div>
                        <div class="side-bar-item">
                            <a href="blog.html" class="side-item-title">web前段优化策略</a>
                            <div class="side-item-info">2.9k阅读 | 一周前</div>
                        </div>
                        <div class="side-bar-item">
                            <a href="blog.html" class="side-item-title">浅析Django项目优化</a>
                            <div class="side-item-info">1.4k阅读 | 两小时前</div>
                        </div>
                        <div class="side-bar-item">
                            <a href="blog.html" class="side-item-title">浅析Django项目优化</a>
                            <div class="side-item-info">4.1k阅读 | 两天前</div>
                        </div>

                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<div class="modal-footer">
    <address class="text-center">
        <p>关于博客园 | 联系我们 | 广告服务 | ©2004-2018博客</p>
        <div href="#">first.last@example.com</div>
    </address>
</div>
</body>
</html>
{{end}}
