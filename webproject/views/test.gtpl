<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>VueTest</title>

		
<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.19.2/axios.js"></script>
</head>
<body>
<div id="app">



	  <div>
	  <button onclick="doSend(this)" id="/blogadd">巴拉拉小魔仙，点击就送</button>
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
				<button id="/blogadd" class="btn btn-primary" onclick="doSend1(this)">添加</button>
			</form>
		</main>
	</div>

</div>

<script>
 function doSend(e){
		console.log("send")
		var js1= $("#form111").serializeArray();
		var data1 = JSON.stringify(transformToJson(js1));
		console.log(data)
		axios({
			headers: {
				'Content-Type': 'application/json'
			},
		method: 'post',
		url: e.id,
		data: data1
		})
		.then(function (res){
			console.log("hearders:",res.headers);
			console.log("status:",res.status);
			console.log("res",res.data.msg)
			alert(res.data.msg);
			window.location.reload();
		})
		.catch(function (err){
			console.log("post err",data);
			alert(data);
			window.location.reload();
		})
	}
</script>

</body>
</html>
