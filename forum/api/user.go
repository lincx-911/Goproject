package api

import (
	"fmt"
	"forum/common"
	"forum/model"
	"forum/service"
	"io"
	
	"os"
	"strconv"


	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/gomail.v2"
)

//GetUserByIDHandle 通过params传输id查找用户
func GetUserByIDHandle(ctx *gin.Context){
	id,err:=strconv.Atoi(ctx.Params.ByName("id"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	user,err:=service.GetUerByIDService(id)
	if err!=nil{
		common.HandleRecordNotFound(ctx)
	}
	common.HandleOkReturn(ctx,user)
}
//GetUserBySelfHandle 不通过参数获取用户信息
func GetUserBySelfHandle(ctx *gin.Context){
	id := ctx.GetInt("uid")
	if id==0{
		common.HandleParamsError(ctx)
		return
	}
	user,err:=service.GetUerByIDService(id)
	if err!=nil{
		common.HandleRecordNotFound(ctx)
	}
	common.HandleOkReturn(ctx,user)
}
//UpdateUserHandle 更新用户
func UpdateUserHandle(ctx *gin.Context)  {
	var user model.User
	_ = ctx.BindJSON(&user)
	fmt.Println("user",user)
	if user.ID==0||user.Username==""{
		common.HandleParamsError(ctx)
		return 
	}
	fmt.Println("user",user)
	if ok,err:=service.UpdateUserService(&user);!ok{
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}

//DeleteUserHandle 删除用户
func DeleteUserHandle(ctx *gin.Context){
	id,err:=strconv.Atoi(ctx.Params.ByName("id"))
	if err!=nil{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.DeleteUserService(id);!ok{
		if err==gorm.ErrRecordNotFound{
			common.HandleRecordNotFound(ctx)
			return
		}
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}
//UserUpdateImg 保存用户头像
func UserUpdateImg(ctx *gin.Context) {
	fmt.Println("进来了img")
	uid := ctx.GetInt("uid")
	fmt.Println(uid)
	if uid == 0{
		common.HandleParamsError(ctx)
		return
	}
	id:=strconv.Itoa(uid)
	fmt.Println(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	file,header,err:=ctx.Request.FormFile("file")
	dir := "static/img/user/"+id
	exit,err:=PathExists(dir)
	if err!=nil{
		common.HandleServerError(ctx,err)
		return
	}
	if !exit{
		err = os.Mkdir(dir, os.ModePerm)
		if err!=nil{
			fmt.Println("mkdir error",err)
			common.HandleOperationError(ctx,err)
			return
		}
	}
	filename:=header.Filename
	out,err:=os.Create(dir+"/"+filename)
	defer out.Close()
	if err!=nil{
		fmt.Println("out error",err)
		common.HandleServerError(ctx,err)
		return
	}
	_,err=io.Copy(out,file)
	if err!=nil{
		fmt.Println("iocopy error",err)
		common.HandleServerError(ctx,err)
		return
	}
	relativeurl := "/img/user/" + id + "/" + filename
	if ok,err:=service.UpdateUserImgService(uid,relativeurl);!ok{
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,relativeurl)
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
//Sendemail 发送邮件
func Sendemail(ctx *gin.Context) {

	var mailConf model.MailboxConf
	
	email:=ctx.Query("email")
	if len(email) == 0 {
		common.HandleParamsError(ctx)
		return
	}
	mailConf.Title = "邮箱验证"
	mailConf.Body = common.GenValidateCode(6)
	mailConf.RecipientList = []string{email}
	mailConf.Sender = `495572661@qq.com`
	mailConf.SPassword = "jxkydqqkgoqfbhhg"
	mailConf.SMTPAddr = `smtp.qq.com`
	mailConf.SMTPPort = 25

	m := gomail.NewMessage()
	m.SetHeader(`From`, mailConf.Sender)
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, mailConf.Body)
	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	if err != nil {
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,mailConf.Body)
}

//UpdatePsw 更改密码
func UpdatePsw(ctx *gin.Context){

	email:=ctx.Query("email")
	password:=ctx.Query("password")
	if len(email) == 0 || password==""{
		common.HandleParamsError(ctx)
		return
	}
	if ok,err:=service.UpdatePassword(email,password);!ok{
		common.HandleServerError(ctx,err)
		return
	}
	common.HandleOkReturn(ctx,nil)
}

