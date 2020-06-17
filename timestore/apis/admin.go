package apis

import (
	"fmt"
	m "miniprogram/models"
	jwt "miniprogram/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

//AdminLogin 管理员登录
func AdminLogin(c *gin.Context) {
	var admin m.Admin
	admin.Aname = c.Request.FormValue("name")
	admin.Password = c.Request.FormValue("password")
	admin, err := admin.SelectAdminbyName()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		tokenString, _ := jwt.GenToken(admin.Aname)
		c.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"token": tokenString,
		})
	}

}

//AdminAdd 添加管理员
func AdminAdd(c *gin.Context) {
	var admin m.Admin
	c.BindJSON(&admin)
	fmt.Println(admin)
	id, err := m.InsertAdmin(admin)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		admin.Aid = id
		c.JSON(http.StatusOK, gin.H{
			"data": admin,
		})
	}
}

//AdminDel 删除管理员
func AdminDel(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	fmt.Println("id=", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	var admin m.Admin
	admin, err = m.DeleteAdmin(id)
	if admin.Aid == 0 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err,
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "删除成功",
			"data": admin,
		})
	}
}

//GetallAdmins 管理员列表
func GetallAdmins(c *gin.Context) {
	admins, err := m.GetallAdmins()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  err,
			"data": "",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": admins,
		})
	}
}

//UpdateAdmin 更新管理员信息
func UpdateAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}
	var admin m.Admin
	admin, err = m.SelectAdminbyID(id)
	if admin.Aid == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "admin not found"})

	} else {
		_ = c.BindJSON(&admin)
		admin.Aid = id
		_, err = m.UpdateAdmin(admin)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": "update admin failed"})
			return
		}
		c.JSON(http.StatusOK, admin)
	}
}
