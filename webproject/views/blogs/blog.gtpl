{{define "blog"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>index</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="statics/css/bootstrap.min1.css">
    <link rel="stylesheet" href="statics/css/main.css">
    <script type="text/javascript">
        function send(){
            console.log(localStorage.token);
            var currentCookie=document.cookie.split(";")[0];
            $.ajax({
            //几个参数需要注意一下
                type: "GET",//方法类型
                dataType: "json",//预期服务器返回的数据类型
                url: "/blogindex" ,//url
                headers:{
                    'Content-Type':'application/json',
                    'Token':currentCookie
                },
                // beforeSend:function(request){
                //     request.setRequestHeader("Token",localStorage.token);
                // }
            });
        }
        window.onload = send;
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
            <div class="col-sm-offset-2">
                <ul class="nav navbar-nav navbar-list">
                    <li><a href="#">全部分类</a></li>
                    <li><a href="#">语言</a></li>
                    <li><a href="#">数据结构</a></li>
                    <li><a href="#">网络</a></li>
                    <li><a href="#">操作系统</a></li>
                    <li><a href="#">算法</a></li>
                    <li><a href="#">实战</a></li>
                </ul>
            </div>
            <ul class="nav navbar-nav navbar-right">
                <li><a href="/login">登陆</a></li>
                <li><a href="/reginster">注册</a></li>
            </ul>
        </div>
    </div>
</div>

<div class="container">
    <div class="row">
        <div class="col-sm-2 hidden-xs">
            <div class="list-group side-bar">
            
                <a class="list-group-item active" role="button">随笔</a>
                {{range .tags}}
                <a class="list-group-item" role="button">{{.}}</a>
                {{end}}
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
                        <a href="#" class="title">
                            记录点滴，记录成长。用博客记录技术与经验的积累，在这里找到志同道合的朋友，编程的乐趣。
                        </a>
                        <div class="info">
                            <span class="avatar"><img src="statics/image/v2-2d45613b0fb8cdd36f53c3b31d0c6ee8_hd.jpg"
                                                      alt="avatar"></span>
                            <span>散人</span> |
                            <span>2.8K</span>
                            <span class="glyphicon glyphicon-thumbs-up blog-hot" aria-hidden="true"></span> |
                            <span>5分钟前</span>
                        </div>
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