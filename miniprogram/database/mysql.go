package database
import (
	"fmt"
	"github.com/jinzhu/gorm"
	//"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"//加载数据库
  )
//DB 数据库操作
var DB *gorm.DB

func init()  {
	var err error
	DB,err=gorm.Open("mysql","root:123456@tcp(8.129.212.77:3306)/mini")
	
	if err != nil{
		fmt.Printf("mysql connect error %v",err)
	}
	fmt.Println("数据库连接成功")
}