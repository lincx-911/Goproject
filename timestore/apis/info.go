package apis

import (
	"fmt"
	m "miniprogram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//IssueInfo 发布招募信息
func IssueInfo(c *gin.Context) {
	var info m.Info
	_ = c.BindJSON(&info)
	fmt.Println("招募队友信息", info)
	_, err := m.SelectUserbyID(info.UID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
		return
	}
	info.Rid, err = m.InsertReinfo(info)
	if err != nil {
		c.JSON(500, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "发布成功",
			"data": info,
		})
	}
}

//AdditInfo 编辑招募信息
func AdditInfo(c *gin.Context) {
	var confo m.Info
	c.BindJSON(&confo)
	fmt.Println("confo", confo)
	cf, err := m.SelectReinfobyID(confo.Rid)
	fmt.Println("cf", cf)
	if cf.Rid == 0 || err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
		return
	}
	fmt.Println("update前", cf)
	_, err = m.UpdateReinfo(confo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "服务器错误",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "更新成功",
			"data": confo,
		})
	}
}

//DeleteInfo 删除招募信息
func DeleteInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	fmt.Println("id=", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
		return
	}
	_, err = m.DeleteReinfo(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "请求失败",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "删除成功",
			"data": "",
		})
	}
}

//GetinfobyID 通过id查询招募信息
func GetinfobyID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	fmt.Println("id=", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求错误",
			"data": "",
		})
		return
	}
	coninfo, err := m.SelectReinfobyID(id)
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

//Getcontestfromtag 通过标签获取招募信息
func Getcontestfromtag(c *gin.Context) {
	var confo m.Info
	c.BindJSON(&confo)
	coninfos, err := m.SelectReinfoByTag(confo.Tag)
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

//Getcontests 获取全部招募信息
func Getcontests(c *gin.Context) {
	coninfos, err := m.GetallReinfos()
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
//Getfreecontests 获取全部时间换时间信息
func Getfreecontests(c *gin.Context) {
	var confo m.Info
	c.BindJSON(&confo)
	coninfos, err := m.SelectReinfoByTtot(confo.Ttot)
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
