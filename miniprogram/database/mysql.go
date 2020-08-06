package database
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"//加载数据库
  )
//DB 数据库操作
var DB *gorm.DB

func init()  {
	var err error
	DB,err:=gorm.Open("mysql","root:root@tcp(8.129.212.77:3306)/mini?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	if err != nil{
		fmt.Printf("mysql connect error %v",err)
	}
	if DB.Error!=nil{
		fmt.Printf("database error %v",DB.Error)
	}
}