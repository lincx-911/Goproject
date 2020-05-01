{{define "article"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>blog</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="statics/css/bootstrap.min.css">
    <link rel="stylesheet" href="statics/css/main.css">
</head>
<body>
<div class="navbar navbar-default">
    <div class="container">
        <div class="navbar-header">
            <a class="navbar-brand" href="/blogindex">My Blog</a>
        </div>
        <label class="toggle-label visible-xs-inline-block" for="toggle-checkbox">MENU</label>
        <input class="hidden" id="toggle-checkbox" type="checkbox">
        
    </div>
</div>

<div class="container">
    <div class="col-sm-8">
        <h1 class="blog-title">{{.Title}}</h1>
        <div class="blog-info">
            <span class="avatar"><img src="statics/image/v2-2d45613b0fb8cdd36f53c3b31d0c6ee8_hd.jpg" alt="avatar"></span>
            <span>{{.Author}}</span> |
            <span>{{.Tag}}</span> |
            <span>{{.Date}}</span>
            <label class="label label-info">编程</label>
            <label class="label label-warning">博客</label>
            <label class="label label-success">Java</label>
        </div>
        <div class="blog-content">
            <blockquote>
                <p>博客生活，记录点滴</p>
            </blockquote>
            <img src="statics/image/blog.png">
            {{.Context}}
        </div>
    </div>
    
</div>
</body>
</html>
{{end}}