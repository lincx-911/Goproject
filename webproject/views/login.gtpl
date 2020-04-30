<!DOCTYPE html>
<html lang="en">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<meta name="description" content="">
		<meta name="author" content="">
		<title>Signin Template for Bootstrap</title>
		<!-- Bootstrap core CSS -->
		<link href="statics/css/bootstrap.min.css" rel="stylesheet">
		<!-- Custom styles for this template -->
		<link href="statics/css/signin.css"  rel="stylesheet">
		<script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
		<script type="text/javascript">
		
			function login() {
				$.ajax({
				//几个参数需要注意一下
					type: "POST",//方法类型
					dataType: "json",//预期服务器返回的数据类型
					url: "/login" ,//url
					data: $('#form1').serialize(),
					success: function (result) {w
						console.log(result);//打印服务端返回的数据(调试用)
						const token = result.data.token
						document.cookie=result.data.token
						if (token) {
							alert("login seccess");
							window.location.href="/blogindex";
						}
					},
					error : function() {
						alert("异常！");
					}
				});
			}
			function getSubPost(para, obj, callback) {
				const IP = Base.prefixUrl + para;
				axios.post(IP, obj, { headers: { Authorization: `Bearer ${Base.token}` } }).then((res) => {
					callback(res.data);
				}).catch((error) => {
					if (error.response.status != '504') {
					callback({ message: `请求接口报错！错误码：${error.response.status}` });
					} else {
					callback({ message: '请求超时，请重试。' });
					}
				});
			}
		</script>
	</head>
	<body class="text-center">
		<form class="form-signin" action="/login" method="post" id="form1">
			<img class="mb-4" src="statics/imgage/bootstrap-solid.svg" alt="" width="72" height="72">
			<h1 class="h3 mb-3 font-weight-normal" text="{{.Tip}}">Please sign in</h1>
			<!--判断-->
			{{if .}}
			<p style="color: red">{{.}}</p>
			{{end}}
			<label class="sr-only" >Username</label>
			<input type="text"  name="username" class="form-control" placeholder="Username" required="" autofocus="">
			<label class="sr-only" >Password</label>
			<input type="password" name="password" class="form-control" placeholder="Password" required="">
			<div class="checkbox mb-3">
				<label>
          			<input type="checkbox" value="remember-me"/> 记住我
        </label>
			</div>
			<button class="btn btn-lg btn-primary btn-block" type="submit" onclick="login()">Sign in</button>
			<p class="mt-5 mb-3 text-muted">© 2017-2018</p>
			<a class="btn btn-sm" th:href="@{/index.html(l='zh_CN')}">中文</a>
			<a class="btn btn-sm" th:href="@{/index.html(l='en_US')}">English</a>
		</form>
	</body>

</html>