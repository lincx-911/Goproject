<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
{{define "leftbar"}}
    <div class="col-sm-2 hidden-xs">
            <div class="list-group side-bar">
                <a class="list-group-item active" role="button">随笔</a>
                <a class="list-group-item" role="button">随笔</a>
                <a class="list-group-item" role="button">随笔</a>
                <a class="list-group-item" role="button">随笔</a>
                <a class="list-group-item" role="button">随笔</a>
                <a class="list-group-item" role="button">随笔</a>
            </div>
        </div>
{{end}}
{{define "rightbar"}}
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
{{end}}
</body>
</html>