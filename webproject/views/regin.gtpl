<!DOCTYPE html>
<html>
<header>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
<script src="http://code.jquery.com/jquery-latest.js"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.19.2/axios.js"></script>
</header>
<body>
<div class="container">
    <div class="col-md-6 col-md-offset-3">
        <h2>用户注册</h2>
            <div class="form-group has-feedback">
                <label for="username">用户名</label>
                <div class="input-group">
                    <span class="input-group-addon"><span class="glyphicon glyphicon-user"></span></span>
                    <input id="username" class="form-control" placeholder="请输入用户名" maxlength="20" type="text" onblur="valiuser()">
                </div>
                <span id="infouser" style="color:red;" class="tips"></span>
                <span style="display: none;" class=" glyphicon glyphicon-remove form-control-feedback"></span>
                <span style="display: none;" class="glyphicon glyphicon-ok form-control-feedback"></span>
            </div>

            <div class="form-group has-feedback">
                <label for="password">密码</label>
                <div class="input-group">
                    <span class="input-group-addon"><span class="glyphicon glyphicon-lock"></span></span>
                    <input id="password" class="form-control" placeholder="请输入密码" maxlength="20" type="password" onblur="valipass()">
                </div>

                <span id="infopass" style="color:red;" class="tips"></span>
                <span style="display: none;" class="glyphicon glyphicon-remove form-control-feedback"></span>
                <span style="display: none;" class="glyphicon glyphicon-ok form-control-feedback"></span>
            </div>

            <div class="form-group has-feedback">
                <label for="passwordConfirm">确认密码</label>
                <div class="input-group">
                    <span class="input-group-addon"><span class="glyphicon glyphicon-lock"></span></span>
                    <input id="passwordConfirm" class="form-control" placeholder="请再次输入密码" maxlength="20" type="password" onblur="valipasscom()">
                </div>
                <span id="infopasscom" style="color:red;" class="tips"></span>
                <span style="display: none;" class="glyphicon glyphicon-remove form-control-feedback"></span>
                <span style="display: none;" class="glyphicon glyphicon-ok form-control-feedback"></span>
            </div>
            <!-- <div class="row">
                <div class="col-xs-7">
                    <div class="form-group has-feedback">
                        <label for="idcode-btn">验证码</label>
                        <div class="input-group">
                            <span class="input-group-addon"><span class="glyphicon glyphicon-qrcode"></span></span>
                            <input id="idcode-btn" class="form-control" placeholder="请输入验证码" maxlength="4" type="text">
                        </div>
                        <span style="color:red;display: none;" class="tips"></span>
                        <span style="display: none;" class="glyphicon glyphicon-remove form-control-feedback"></span>
                        <span style="display: none;" class="glyphicon glyphicon-ok form-control-feedback"></span>
                    </div>
                </div>
                <div class="col-xs-5" style="padding-top: 30px">
                    <div id="idcode" style="background: transparent;"></div>
                </div>
            </div> -->

            <div class="form-group has-feedback">
                <label for="email">email</label>
                <div class="input-group">
                    <span class="input-group-addon"><span class="glyphicon glyphicon-envelope"></span></span>
                    <input id="email" class="form-control" placeholder="请输入email" type="email" onblur="valiemail()">
                </div>
                <span  id="infoemail" style="color:red;" class="tips"></span>
                <span style="display: none;" class="glyphicon glyphicon-remove form-control-feedback"></span>
                <span style="display: none;" class="glyphicon glyphicon-ok form-control-feedback"></span>
            </div>

            <!-- <div class="row">
                <div class="col-xs-7">
                    <div class="form-group has-feedback">
                        <label for="idcode-btn">校验码</label>
                        <div class="input-group">
                            <span class="input-group-addon"><span class="glyphicon glyphicon-qrcode"></span></span>
                            <input id="idcode-btn" class="form-control" placeholder="请输入校验码" maxlength="6" type="text">
                        </div>
                        <span style="color:red;display: none;" class="tips"></span>
                        <span style="display: none;" class="glyphicon glyphicon-remove form-control-feedback"></span>
                        <span style="display: none;" class="glyphicon glyphicon-ok form-control-feedback"></span>
                    </div>
                </div> -->
                <!-- <div class="col-xs-5 text-center" style="padding-top: 26px">
                    <button type="button" id="loadingButton" class="btn btn-primary" autocomplete="off">获取邮箱校验码</button>
                </div> -->
            <!-- </div> -->

            <div class="form-group">
                <input class="btn btn-primary" id="submit" value="注册" type="submit" onclick="dosend()">
                <input value="重置" id="reset" class="btn btn-danger" type="reset">
            </div>

            <div class="form-group">
                
            </div>
    </div>
</div>
<script >
    function valiuser(){
        var username = document.getElementById("username").value;
        var info = document.getElementById("infouser");
        if (username==""){
            info.innerHTML='username不能为空';
            // 阻止程序向下执行
            return;
        }else{
            info.innerHTML='';
        }
    }
    function valipass(){
        var password = document.getElementById('password').value;
        var info = document.getElementById('infopass');
        if (password==""){
            info.innerHTML='password不能为空';
            // 阻止程序向下执行
            return;
        }else{
            info.innerHTML='';
        }
    }
    function valipasscom(){
        var passwordcom = document.getElementById('passwordConfirm').value;
        var info = document.getElementById('infopasscom');
        if (passwordcom==""){
            info.innerHTML='password不能为空';
            // 阻止程序向下执行
            return;
        }else{
            info.innerHTML='';
        }
    }
    function valiemail(){
        var email = document.getElementById('email').value;
        var info = document.getElementById('infoemail');
        if (email==""){
            info.innerHTML='email不能为空';
            // 阻止程序向下执行
            return;
        }else{
            info.innerHTML='';
        }
    }
    function dosend(){
        var username = document.getElementById("username").value;
        var password = document.getElementById("password").value;
        var passwordConfirm = document.getElementById("passwordConfirm").value;
        var email = document.getElementById("email").value;
        var data1={
            "username":username,
            "password":password,
            "email":email
        };
        axios({
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'post',
            url: "/reginster",
            data: JSON.stringify(data1)
            })
            .then(function (res){
                alert(res.data.msg);
                console.log(res.headers);
                console.log(res.data);
                console.log(res.status);
                window.location.href="/login";
            })
            .catch(function (err){
                console.log("post err",err);
                alert(err.msg);
                window.location.reload();
            })
    }
</script>
</body>
</html>