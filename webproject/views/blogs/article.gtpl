{{define "article"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>blog</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
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