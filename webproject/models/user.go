package models

import (

	
	"database/sql"
	"fmt"
	"webproject/utils"
	//启动mysql
	_ "github.com/go-sql-driver/mysql"
)

//User 用户模型
type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email string `json:"email" db:"email"`
}
//Getlink 连接数据库
func Getlink() *sql.DB {
	db, err := sql.Open("mysql", "root:xxxxx@tcp(127.0.0.1)/jdbc")//根据实际来改，这里我就不放我的数据库地址了
	utils.CheckError(err)
	return db
}

//Insert 插入用户
func Insert(u User) (users User,err error) {
	db:=Getlink()
	defer db.Close()
	stmt,err:=db.Prepare("INSERT Bloguser SET username=?,password=?,email=?")
	utils.CheckError(err)
	_ ,err=stmt.Exec(u.Username,u.Password,u.Email)
	utils.CheckError(err)
	return u,nil
}
//DeleteUser 删除用户
func DeleteUser(id int) (err error) {
	db:=Getlink()
	defer db.Close()
	fmt.Println("进来了删除user")
	stmt,err := db.Prepare("DELETE FROM Bloguser WHERE id=?")
	if err!=nil{
		return err
	}
	_,err=stmt.Exec(id)
	if err!=nil{
		return err
	}
	db.Exec("ALTER TABLE Bloguser AUTO_INCREMENT = 1")
	fmt.Println("好像成功删除user了？")
	return nil
}
//Update 修改用户
func Update(u User) (err error)  {
	db:=Getlink()
	defer db.Close()
	stmt,err := db.Prepare("UPDATE Bloguser SET username=?,password=?,email=? where id=?")
	utils.CheckError(err)
	_,err=stmt.Exec(u.Username,u.Password,u.Email)
	utils.CheckError(err)
	return nil
}
//GetUserbyName 通过姓名查询用户
func GetUserbyName(name string) (u User,err error){
	db:=Getlink()
	defer db.Close()
	var user User
	rows,err:=db.Query("SELECT * FROM Bloguser WHERE username=?",name)
	if err!=nil{
		fmt.Println("找不到")
		fmt.Println(err,user)
		return
	}
	for rows.Next(){
		err=rows.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
		if err!=nil{
			fmt.Println("找不到ss")
			fmt.Println(err,user)
			return
		}
	}
	return user,nil
}
//GetUserbyID 通过id查询用户
func GetUserbyID(id int) (u User,err error){
	db:=Getlink()
	defer db.Close()
	var user User
	err=db.QueryRow("SELECT id,username,password,email FROM Bloguser WHERE id=?",id).Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	utils.CheckError(err)
	return user,nil
}
//GetAllUser 得到所有用户
func GetAllUser()(u []User,err error){
	db:=Getlink()
	defer db.Close()
	var users []User
	var user User
	rows,err:=db.Query("SELECT * FROM Bloguser")
	utils.CheckError(err)
	defer rows.Close()
	for rows.Next(){
		rows.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
		utils.CheckError(err)
		users = append(users,user)
	}
	return users,nil
}