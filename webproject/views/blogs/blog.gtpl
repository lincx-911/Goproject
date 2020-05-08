
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
    <link rel="stylesheet" href="statics/css/bootstrap.min.css"><!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <link rel="stylesheet" href="statics/css/main.css">
</head>
<body>


<div class="container">
<nav class="navbar navbar-expand-sm bg-light navbar-light">
  <ul class="nav nav-tabs">
    <li class="nav-item ">
      <a class="nav-link" href="#">全部标签</a>
    </li>
    {{range $k,$v:=.Tags}}
    <li class="nav-item">
      <a class="nav-link" href="#"><span class="menu-title">{{$k}}</span></a>
    </li>
    {{end}}
  </ul>
</nav>
    <div class="row">
        <div class="col-sm-2 hidden-xs">
            <div class="list-group side-bar">

               <p>全部分类</p>
                <ul class="nav navbar-nav">
                    
                    {{range $k,$v:=.Categories}}
                    <li><a href="#">{{$k}}</a>
                    </li>
                    {{end}}
                </ul>
            </div>
        </div>

        <div class="col-sm-7">
            <div class="blog-list">
                <div class="blog-list-item clearfix">
                    <div class="col-xs-7">
                        
                         {{range .Blogs}}
                        <a href="/blogarticle?id={{.ID}}" class="title">
                            {{.Title}}
                        </a>
                         <img src="statics/image/blog.png">
                        <div class="info">
                            <span class="avatar"><img src="statics/image/blog.png"
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
            <div class="side-bar-card side-bar-recommend clearfix">
                <div class="side-bar-title">推荐阅读</div>
                <div class="side-bar-body">
                    <div class="side-bar-list">
                    {{range .Blogs}}
                        {{if le .ID 5}}
                        <div class="side-bar-item">
                            <a href="blog.html" class="side-item-title">{{.Title}}</a>
                            <div class="side-item-info">{{.Tag}} | {{.Date}}</div>
                        </div>
                        {{end}}
                    {{end}}
                        
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
     $(function(){

        $("#sideMenu").metisMenu();   

     })
</script>
</body>
</html>
{{end}}
