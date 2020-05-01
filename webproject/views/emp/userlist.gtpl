{{define "userlist"}}
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
		<link href="statics/css/bootstrap.min.css" rel="stylesheet">

		<!-- Custom styles for this template -->
		<link href="statics/css/dashboard.css" rel="stylesheet">
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
				<main role="main" class="col-md-9 ml-sm-auto col-lg-10 pt-3 px-4">
					<h2>Blog用户</h2>
					<!-- <h2><a class="btn btn-sm btn-success" href="/blogadd">user添加</a></h2> -->
					<div class="table-responsive">
						<table class="table table-striped table-sm">
							<thead>
								<tr>
									<th>id</th>
									<th>username</th>
									<th>email</th>
									<th>操作</th>
								</tr>
							</thead>
							<tbody>
								{{range .}}
								<tr>
								
									<td >{{.ID}}</td>
									<td>{{.Username}}</td>
									<td>{{.Email}}</td>
									<td>
										<!-- <button id="{{.ID}}" del_uri="/blogdelete?id={{.ID}}" class="btn btn-sm btn-danger deleteBtn">删除</button> -->
										<button id="{{.ID}}" onclick="doSend1(this)" class="btn btn-sm btn-danger deleteBtn">删除</button>
									</td>
								</tr>
								{{end}}
							</tbody>
							
						</table>
						<div>
							{{if .}}
							{{else}}
								<h2>好像啥也没有哦</h2>
							{{end}}
						</div>
					</div>
				</main>
					
		</div>

		<!-- Bootstrap core JavaScript
    ================================================== -->
		<!-- Placed at the end of the document so the pages load faster -->
		<script type="text/javascript" src="statics/js/jquery-3.2.1.slim.min.js" th:src="@{/webjars/jquery/3.3.1/jquery.js}"></script>
		<script type="text/javascript" src="statics/js/popper.min.js" th:src="@{/webjars/popper.js/1.11.1/dist/popper.js}"></script>
		<script type="text/javascript" src="statics/js/bootstrap.min.js" th:src="@{/webjars/bootstrap/4.0.0/js/bootstrap.js}"></script>

		<!-- Icons -->
		<script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
		<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
		<script type="text/javascript" src="statics/js/feather.min.js" th:src="@{/asserts/js/feather.min.js}"></script>
		<script>
			feather.replace()
		</script>
		<script>
			function doSend1(e){
				console.log("send")
				var id = e.id
				console.log(id)
				data1={"id":id}
				axios({
					headers: {
						'Content-Type': 'application/json'
					},
					method: 'post',
					url: "/userdelete",
					data: JSON.stringify(data1)
				})
				.then(function (res){
                alert(res.data.msg);
                console.log(res.headers);
                console.log(res.data);
                console.log(res.status);
                window.location.reload();
            	})
				.catch(function (err){
					console.log("post err",err);
					alert(err);
				})
			}
		</script>
	</body>
</html>
{{end}}