{{define "atticle"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>blog</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="statics/css/bootstrap.min1.css">
    <link rel="stylesheet" href="statics/css/main.css">
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
                <li><a href="login.html">登陆</a></li>
                <li><a href="login.html">注册</a></li>
            </ul>
        </div>
    </div>
</div>

<div class="container">
    <div class="col-sm-8">
        <h1 class="blog-title">记录点滴，记录成长。用博客记录技术与经验的积累，在这里找到志同道合的朋友，编程的乐趣。</h1>
        <div class="blog-info">
            <span class="avatar"><img src="statics/image/v2-2d45613b0fb8cdd36f53c3b31d0c6ee8_hd.jpg" alt="avatar"></span>
            <span>散人</span> |
            <span>2.8K热度</span> |
            <span>5分钟前</span>
            <label class="label label-info">编程</label>
            <label class="label label-warning">博客</label>
            <label class="label label-success">Java</label>
        </div>
        <div class="blog-content">
            <blockquote>
                <p>博客生活，记录点滴</p>
            </blockquote>
            <img src="statics/image/blog.png">
            如果你有大量的设置为 inline 属性的标签全部放在一个较窄的容器元素内，在页面上展示这些标签就会出现问题，每个标签就会有自己的一个 inline-block 元素（就像图标一样）。
            解决的办法是为每个标签都设置为 display: inline-block; 属性。
            如果你有大量的设置为 inline 属性的标签全部放在一个较窄的容器元素内，在页面上展示这些标签就会出现问题，每个标签就会有自己的一个 inline-block 元素（就像图标一样）。
            解决的办法是为每个标签都设置为 display: inline-block; 属性。
        </div>
    </div>
    <div class="col-sm-4 hidden-xs">
        {{template "rightbar"}}
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