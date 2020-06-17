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

//OganizeLogin 组织登录
func OganizeLogin(c *gin.Context) {
	var ogn m.Organize
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	fmt.Println("email:", email)
	fmt.Println("password", password)
	ogn, err := m.Selectorgbyemail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":   "用户不存在",
			"data":  "",
			"token": "",
		})
		return
	}

	if ogn.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "账号或者密码出错",
			"data":  "",
			"token": "",
		})
	} else {
		tokenString, _ := jwt.GenToken(ogn.Oname)
		c.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"data":  ogn,
			"token": tokenString,
		})
	}
}

//OrganizeRegist 组织注册
func OrganizeRegist(c *gin.Context) {
	var org m.Organize
	c.BindJSON(&org)
	id, err := m.InsertOrganize(org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "服务器错误",
			"data":  "",
			"token": "",
		})
	} else {
		org.Oid = id
		c.JSON(http.StatusOK, gin.H{
			"msg":  "注册成功",
			"data": org,
		})
	}
}

//OrganizeDel 删除组织
func OrganizeDel(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	fmt.Println("id=", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	var org m.Organize
	org, err = m.DeleteOrganize(id)
	if org.Oid == 0 || err != nil {
		c.JSON(http.StatusNotFound, gin.H{
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

//OrganizeUpdate 组织信息更新
func OrganizeUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	var org m.Organize
	org, err = m.SelectorgbyID(id)
	if org.Oid == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})

	} else {
		_ = c.BindJSON(&org)
		org.Oid = id
		_, err = m.UpdateOrganize(org)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": "update user failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "更新成功",
			"data": org,
		})
	}

}

//GetOrgbyid 通过id获取组织
func GetOrgbyid(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	org, err := m.SelectorgbyID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "用户不存在",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "查询成功",
			"data": org,
		})
	}
}

//GetallOrganizes 获取全部组织信息
func GetallOrganizes(c *gin.Context) {
	orgs, err := m.GetallOrganizes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "获取失败",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": orgs,
		})
	}
}

//IssueContest 发布赛事信息
func IssueContest(c *gin.Context) {
	var cf m.Contestinfo
	_ = c.BindJSON(&cf)
	fmt.Println("招募队友信息", cf)
	org, err := m.SelectorgbyID(cf.Oid)
	if org.Oid == 0 || err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "组织不存在",
			"data": "",
		})
		return
	}
	cf, err = m.InsertContestinfo(cf)
	if err != nil {
		c.JSON(500, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
	} else {
		fmt.Println("cf", cf)
		c.JSON(http.StatusOK, gin.H{
			"msg":  "发布成功",
			"data": cf,
		})
	}
}

//AdditContest 编辑赛事信息
func AdditContest(c *gin.Context) {
	var confo m.Contestinfo
	c.BindJSON(&confo)
	cf, err := m.SelectConinfobyID(confo.Cid)
	if cf.Cid == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
		return
	}
	fmt.Println("update前", cf)
	_, err = m.UpdateContestinfo(confo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "更新失败",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "更新成功",
			"data": confo,
		})
	}
}

//DeleteContest 删除赛事信息
func DeleteContest(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	fmt.Println("id=", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	_, err = m.DeleteContestinfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
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

//Getcontest 通过id查询赛事
func Getcontest(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	fmt.Println("id=", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	coninfo, err := m.SelectConinfobyID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  err,
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": coninfo,
		})
	}
}

//Getcontestfromtag 通过标签获取赛事信息
func Getcontestfromtag(c *gin.Context) {
	var confo m.Contestinfo
	c.BindJSON(&confo)
	coninfos, err := m.SelectConinfoByTag(confo.Ctag)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  err,
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": coninfos,
		})
	}
}

//GetcontestfromPublic 通过发布方获取赛事信息
func GetcontestfromPublic(c *gin.Context) {
	var confo m.Contestinfo
	c.BindJSON(&confo)
	fmt.Println("publicer:",confo.Publicer)
	coninfos, err := m.SelectConinfoBypublicer(confo.Publicer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  err,
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": coninfos,
		})
	}
}
//Getcontests 获取全部赛事信息
func Getcontests(c *gin.Context){
	coninfos,err:=m.GetallConinfos()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  err,
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": coninfos,
		})
	}
}

//Saveoimg 保存组织头像
func Saveoimg(c *gin.Context) {
	// name:=c.PostForm("name")//根据name属性获取文件名
	// fmt.Println(name)
	id := c.Params.ByName("id")
	oid,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
	}
	oselect,err:=m.SelectorgbyID(oid)
	if oselect.Oid==0||err!=nil{
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
	dir := "static/img/organize/" + id
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
		relativeurl := "/img/organize/" + id + "/" + filename
		
		if err=m.Uploaduserimage(oid,relativeurl);err!=nil{
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
//Uploadconfile 上传证明文件
func Uploadconfile(c *gin.Context){
	id := c.Params.ByName("id")
	cid,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
	}
	cselect,err:=m.SelectConinfobyID(cid)
	if cselect.Cid==0||err!=nil{
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
	fmt.Println("body:")
	fmt.Println("header1:",header)
	fmt.Println(file)
	dir := "static/file/contests/" + id
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
		relativeurl := "/file/contests/" + id + "/" + filename
		if err=m.Uploaduserimage(cid,relativeurl);err!=nil{
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