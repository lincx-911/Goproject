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
		<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
		<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.19.2/axios.js"></script>
		<script type="text/javascript">
		
			// function login() {
			// 	$.ajax({
			// 	//几个参数需要注意一下
			// 		type: "POST",//方法类型
			// 		dataType: "json",//预期服务器返回的数据类型
			// 		url: "/login" ,//url
			// 		data: $('#form1').serialize(),
			// 		success: function (result) {w
			// 			console.log(result);//打印服务端返回的数据(调试用)
			// 			const token = result.data.token
			// 			document.cookie=result.data.token
			// 			if (token) {
			// 				alert("login seccess");
			// 				window.location.href="/blogindex";
			// 			}
			// 		},
			// 		error : function() {
			// 			alert("异常！");
			// 		}
			// 	});
			// }
			function login(){
				var username = document.getElementById("username").value;
				var password = document.getElementById("password").value;
				console.log(username,password);
				data1={
					"username":username,
					"password":password
				}
				console.log(data1)
				axios({
				headers: {
					'Content-Type': 'application/json'
				},
				method: 'post',
				url: "/login",
				data: JSON.stringify(data1)
				})
				.then(function (res){
					console.log(res);
					console.log(res.data.data);
					console.log(res.status);
					// window.localStorage.setItem("Authorization",res.data.data);//保存token
					document.cookie = res.data.data;
					alert(res.data.data);
					window.location.href="/index";
				})
				.catch(function (err){
					console.log("post err",err);
					alert(err);
					window.location.reload();
				})
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
		<div class="form-signin" >
			<img class="mb-4" src="statics/image/bootstrap-solid.svg" alt="" width="72" height="72">
			<h1 class="h3 mb-3 font-weight-normal">登录后台管理</h1>
			<!--判断-->
			<label class="sr-only" >Username</label>
			<input type="text" id="username" name="username" class="form-control" placeholder="Username" required="" autofocus="">
			<label class="sr-only" >Password</label>
			<input type="password" id="password" name="password" class="form-control" placeholder="Password" required="">
			<div class="checkbox mb-3">
				<label>
          			<input type="checkbox" value="remember-me"/> 记住我
        		</label>
			</div>
			<button class="btn btn-lg btn-primary btn-block" type="submit" onclick="login()">登录</button>
			<button class="btn btn-lg btn-primary btn-block" type="submit" onclick="login()">注册</button>
			<p class="mt-5 mb-3 text-muted">© 2017-2018</p>
			<a class="btn btn-sm" th:href="@{/index.html(l='zh_CN')}">中文</a>
			<a class="btn btn-sm" th:href="@{/index.html(l='en_US')}">English</a>
		</div>
	</body>

</html>