package models

import (
	
	//启动mysql
	_ "github.com/go-sql-driver/mysql"
	"webproject/utils"
	"database/sql"
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
	db, err := sql.Open("mysql", "root:123456@tcp(120.78.6.205)/jdbc")
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
//Delete 删除用户
func Delete(u User) (err error) {
	db:=Getlink()
	defer db.Close()
	stmt,err := db.Prepare("DELETE FROM Bloguser WHERE id=?")
	utils.CheckError(err)
	_,err=stmt.Exec(u.ID)
	utils.CheckError(err)
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
	err=db.QueryRow("SELECT id,username,password,email FROM Bloguser WHERE username=?",name).Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	utils.CheckError(err)
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
func GetAllUser(id int)(u []User,err error){
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