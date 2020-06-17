package apis

import (
	"fmt"
	"io"
	m "miniprogram/models"
	jwt "miniprogram/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

//UserLogin 用户登录登录
// @Tags 用户
// @Id 1
// @Summary 用户登录
// @Security ApiKeyAuth
// @Param id path int true "用户id"
// @Param token header string true "用户token"
// @Param name query string false "用户名"
// @Param img formData file false "文件"
// @Success 200 object  response.Result "成功"
// @Failure 400 object  response.Result "失败"
// @Router /api/v1/{id} [get]
func UserLogin(c *gin.Context) {
	var user m.User
	user.Uemail = c.Request.FormValue("email")
	user.Password = c.Request.FormValue("password")
	user, err := user.SelectUserbyemail()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "账号或者密码出错",
			"data":  "",
			"token": "",
		})
	} else {
		tokenString, _ := jwt.GenToken(user.Uname)
		c.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"data":  user,
			"token": tokenString,
		})
	}

}

//UserRegist 用户注册
func UserRegist(c *gin.Context) {
	var user m.User
	c.BindJSON(&user)
	id, err := m.InsertUser(user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		user.UID = id
		c.JSON(http.StatusOK, gin.H{
			"msg":  "注册成功",
			"data": user,
		})
	}
}

//UserDel 删除用户信息
func UserDel(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	fmt.Println("id=", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	var user m.User
	user, err = m.DeleteUser(id)
	if user.UID == 0 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err,
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "删除成功",
			"data": "",
		})
	}
}

//GetallUsers 用户列表
func GetallUsers(c *gin.Context) {
	users, err := m.GetallUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  err,
			"data": "",
		})

	} else {
		if len(users) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "用户信息为空",
				"data": "",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": users,
		})
	}
}

//UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	var user m.User
	user, err = m.SelectUserbyID(id)
	if user.UID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})

	} else {
		_ = c.BindJSON(&user)
		user.UID = id
		_, err = m.UpdateUser(user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": "update user failed"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

//GetmyRecruir 获取个人发布的招募信息
func GetmyRecruir(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	var reinfo m.Reinfo
	reinfo, err = m.SelectReinfobyID(id)
	fmt.Println("reinfo:", reinfo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "获取信息失败",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": reinfo,
		})
	}

}

//IssueRecruit 发布招募信息
func IssueRecruit(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	var rf m.Reinfo
	_ = c.BindJSON(&rf)
	fmt.Println("招募队友信息", rf)
	rf.Rid, err = m.InsertReinfo(rf, id)
	if err != nil {
		c.JSON(500, gin.H{
			"msg":  "发布失败1",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "发布成功",
			"data": rf,
		})
	}

}

//AditRecruit 编辑招募信息
func AditRecruit(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	fmt.Println("select 前")
	var reinfo m.Reinfo
	reinfo, err = m.SelectReinfobyID(id)
	fmt.Println("select 后", reinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
	}
	c.BindJSON(&reinfo)
	reinfo.Rid = id
	fmt.Println("update前", reinfo)
	_, err = m.UpdateReinfo(reinfo)
	fmt.Println("update后")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "招募信息更新失败",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "招募信息更新成功",
			"data": reinfo,
		})
	}
}

//DeleteRecruit 删除招募信息
func DeleteRecruit(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	fmt.Println("id:", id)
	_, err = m.DeleteReinfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "删除失败",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "删除成功",
			"data": "",
		})
	}
}

//AddCollectcon 添加赛事收藏
func AddCollectcon(c *gin.Context) {
	var cl m.Uclloctcon //收藏
	_ = c.BindJSON(&cl)
	fmt.Println("添加收藏信息", cl)
	cl, err := m.InsertUclloctcon(cl)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg":  "发布失败",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "发布成功",
			"data": cl,
		})
	}
}

//AddCollectre 添加招募收藏
func AddCollectre(c *gin.Context) {
	var re m.Uclloctre //收藏
	_ = c.BindJSON(&re)
	fmt.Println("添加收藏信息", re)
	_, err := m.InsertUclloctre(re)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg":  "发布失败",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "发布成功",
			"data": re,
		})
	}
}

//DeleteColectcon 删除赛事收藏
func DeleteColectcon(c *gin.Context) {
	var re1 m.Uclloctcon //收藏re
	_ = c.BindJSON(&re1)
	_, err := m.DelUclloctcon(re1)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "删除成功",
			"data": "",
		})
	}
}

//DeleteColectre 删除招募信息收藏
func DeleteColectre(c *gin.Context) {
	var re m.Uclloctre
	_ = c.BindJSON(&re)
	_, err := m.DelUclloctre(re)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "删除成功",
			"data": "",
		})
	}
}

//GetallCollectcon 获取全部赛事收藏
func GetallCollectcon(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	con, err := m.GetallUclloctcon(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": con,
		})
	}
}

//GetallCollectre 获取全部招募收藏
func GetallCollectre(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	con, err := m.GetallUclloctre(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": con,
		})
	}
}

//Saveuimg 保存用户头像
func Saveuimg(c *gin.Context) {
	// name:=c.PostForm("name")//根据name属性获取文件名
	// fmt.Println(name)
	fmt.Println("进来了")
	id := c.Params.ByName("id")
	uid,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
	}
	uselect,err:=m.SelectUserbyID(uid)
	if uselect.UID==0||err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
	}
	defer c.Request.Body.Close()
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("header:", c.Request.Header)
	fmt.Println("body:", c.Request.Body)
	fmt.Println(file)
	dir := "static/img/user/" + id
	exit, err := PathExists(dir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exit {
		fmt.Printf("no dir![%v]\n", dir)
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("succeed mkdir")
		return
	}
	fmt.Println("exited")
	filename := header.Filename
	fmt.Println("filename:", filename)
	out, err := os.Create(dir + "/" + filename)
	defer out.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

	} else {
		relativeurl := "/img/user/" + id + "/" + filename
		
		if err=m.Uploaduserimage(uid,relativeurl);err!=nil{
			c.JSON(http.StatusNotFound, gin.H{
				"status": 500,
				"msg":    "服务器错误",
				"data": "" ,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "上传成功",
			"data": relativeurl ,
		})

	}
}

//PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
//Uploadrefile 上传赛事文件
