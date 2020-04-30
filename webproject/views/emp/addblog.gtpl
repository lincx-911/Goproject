
{{define "addblog"}}
<!DOCTYPE html>
<!-- saved from url=(0052)http://getbootstrap.com/docs/4.0/examples/dashboard/ -->
<html lang="en">

	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<meta name="description" content="">
		<meta name="author" content="">

		<title>Dashboard Template for Bootstrap</title>
		<!-- Bootstrap core CSS -->
		<link href="statics/css/bootstrap.min.css"rel="stylesheet">

		<!-- Custom styles for this template -->
		<link href="statics/css/dashboard.css"rel="stylesheet">
		<style type="text/css">
			/* Chart.js */
			@-webkit-keyframes chartjs-render-animation {
				from {
					opacity: 0.99
				}
				to {
					opacity: 1
				}
			}
			
			@keyframes chartjs-render-animation {
				from {
					opacity: 0.99
				}
				to {
					opacity: 1
				}
			}
			
			.chartjs-render-monitor {
				-webkit-animation: chartjs-render-animation 0.001s;
				animation: chartjs-render-animation 0.001s;
			}
		</style>
	</head>

	<body>
		{{template "content1"}}
		<div class="container-fluid">
			<div class="row">
				<!--引入侧边栏-->
				{{template "content2"}}
			</div>
		</div>
			<div class="center111">
				<main role="main" class="col-md-9 ml-sm-auto col-lg-10 pt-3 px-4">
					<form id="form111">
						<div class="form-group">
							<label>Tag</label>
							<input name="tag" type="text" class="form-control1" placeholder="golang">
						</div>
						<div class="form-group">
							<label>Categorie</label>
							<input name="categorie" type="text" class="form-control1" placeholder="后端语言" >
						</div>
						<div class="form-group">
							<label>Title</label>
							<input name="title" class="form-control1" placeholder="golang学习">
						</div>
						<div class="form-group">
							<label>Context</label>
							<!-- <textarea name="context" rows="12" cols="30" class="form-control1">此处写下blog内容
							</textarea> -->
							<input name="context" type="text" class="form-control1" placeholder="人生苦短，js搞死人">
						</div>
						<div class="form-group">
							<label>Date</label>
							<input id="today" name="date" type="text" class="form-control1">
						</div>
						<input  name="id" type="hidden" class="form-control1" value="0">
						
					</form>
					<button id="/blogadd" class="btn btn-primary" onclick="doSend1(this)">添加</button>
				</main>
			</div>
	

		<!-- Bootstrap core JavaScript
    ================================================== -->
		<!-- Placed at the end of the document so the pages load faster -->
		<script type="text/javascript" src="statics/js/jquery-3.2.1.slim.min.js"></script>
		<script type="text/javascript" src="statics/js/popper.min.js" ></script>
		<script type="text/javascript" src="statics/js/bootstrap.min.js" ></script>
		<!-- Icons -->
		<script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
		<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
		<script type="text/javascript" src="statics/js/feather.min.js" ></script>
		<script type="text/javascript">
			feather.replace()
			function today(){//构建方法
					var today=new Date();//new 出当前时间
					var h=today.getFullYear();//获取年
					var m=today.getMonth()+1< 10 ? "0" + (today.getMonth() + 1) : today.getMonth() + 1;;//获取月
					var d=today.getDate() < 10 ? "0" + today.getDate() : today.getDate();//获取日
					var H = today.getHours()< 10 ? "0" + today.getHours() : today.getHours();//获取时
					var M = today.getMinutes()< 10 ? "0" + today.getMinutes() : today.getMinutes();//获取分
					var S=today.getSeconds()< 10 ? "0" + today.getSeconds() : today.getSeconds();
					return h+"-"+m+"-"+d+" "+H+":"+M+":"+S; //返回 年-月-日 时:分:秒
			}
			document.getElementById("today").value = today();//将获取到的 年-月-日 时:分:秒 赋给input文本输入框
			function transformToJson(formData){
				var obj={}
				for (var i in formData) {
					obj[formData[i].name]=formData[i]['value'];
				}
				return obj;
			}
			function doSend1(e){
				console.log("send")
				 var js1= $("#form111").serializeArray();
				 var data1 = JSON.stringify(transformToJson(js1));
				 var url=e.id
				 console.log(data1)
				 alert(data1)
				axios({
					headers: {
						'Content-Type': 'application/json'
					},
				method: 'post',
				url: url,
				data: data1
				})
				.then(function (res){
					alert(res.data.msg)
					console.log(res.headers);
					console.log(res.data);
					console.log(res.status);
					alert(res.data.msg);
					window.location.href="/bloglist";
				})
				.catch(function (err){
					console.log("post err",err);
					alert(err);
					window.location.href="/bloglist";
				})
			}
    </script>
	</body>
</html>
{{end}}