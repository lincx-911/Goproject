package respository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //加载数据库
)

//DB 数据库操作
var DB *gorm.DB

func init()  {
	var err error
	DB,err=gorm.Open("mysql","root:123456@tcp(120.78.6.205:3306)/forum?charset=utf8")
	if err!=nil{
		log.Fatalf("mysql connext error %v",err)
	}
	//DB.AutoMigrate(&model.UserClike{})
	log.Println("数据库连接成功")
}