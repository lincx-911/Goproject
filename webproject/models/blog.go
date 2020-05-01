package models

import (
	"database/sql/driver"
	"time"
	"webproject/utils"

	//启动mysql
	_ "github.com/go-sql-driver/mysql"
)

//TimeFormat 时间格式
const TimeFormat = "2006-01-02 15:04:05"

//LocalTime 本地时间
type LocalTime time.Time

//Blog 模型
type Blog struct {
	ID        string    `json:"id" db:"id"`
	Tag       string    `json:"tag" db:"tag"`
	Categorie string    `json:"categorie" db:"categorie"`
	Title     string    `json:"title" db:"title"`
	Author    string    `json:"author" db:"author"`
	Context   string    `json:"context" db:"context"`
	Date      LocalTime `json:"date" db:"date"`
}

//Tags 标签模型
type Tags struct {
	Tag string `json:"tag"`
}

//Categories 分类
type Categories struct {
	Categorie string `json:"catagorie"`
}

//UnmarshalJSON 解析时间json
func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}
	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

//MarshalJSON 时间转json
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

//Value 写入 mysql 时调用
func (t LocalTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

//Scan 检出 mysql 时调用
func (t *LocalTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	timestring := v.(string)
	tTime, _ := time.ParseInLocation("2006-01-02 15:04:05 +0800 CST", timestring, time.Local)
	*t = LocalTime(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

// Getlink 连接数据库
// func getlink() *sql.DB {
// 	db, err := sql.Open("mysql", "root:123456@tcp(120.78.6.205)/jdbc")
// 	utils.CheckError(err)
// 	return db
// }

//InsertBg 插入blog
func InsertBg(blog Blog) (err error) {
	db := Getlink()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO BlogArticle (tag,categorie,title,context,author,date) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(blog.Tag, blog.Categorie, blog.Title, blog.Context, blog.Author, blog.Date)
	return err
}

//DeleteBg 删除blog
func DeleteBg(id string) (err error) {
	db := Getlink()
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM BlogArticle WHERE id=?")
	utils.CheckError(err)
	_, err = stmt.Exec(id)
	utils.CheckError(err)
	db.Exec("ALTER TABLE BlogArticle AUTO_INCREMENT = 1")
	return nil
}
//UpdateBg 修改blog
func UpdateBg(blog Blog) (err error) {
	db := Getlink()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE BlogArticle SET tag=?,categorie=?,title=?,context=?,author=?,date=? WHERE id=?")
	utils.CheckError(err)
	_, err = stmt.Exec(blog.Tag, blog.Categorie, blog.Title, blog.Context, blog.Author, blog.Date, blog.ID)
	utils.CheckError(err)
	return nil
}

//GetBlogbyID 通过id查询文章
func GetBlogbyID(id int) (b Blog, err error) {
	db := Getlink()
	defer db.Close()
	var blog Blog
	var blogtime string
	err = db.QueryRow("SELECT id,tag,categorie,title,context,author,date FROM BlogArticle WHERE id=?", id).Scan(&blog.ID, &blog.Tag, &blog.Categorie, &blog.Title, &blog.Context, &blog.Author, &blogtime)
	utils.CheckError(err)
	time1, _ := time.ParseInLocation("2006-01-02 15:04:05", blogtime, time.Local)
	blog.Date = LocalTime(time1)
	return blog, nil
}

//GetAllBlog 得到所有文章
func GetAllBlog() (b []Blog, err error) {
	db := Getlink()
	defer db.Close()
	var blogs []Blog
	var blog Blog
	var blogtime string
	rows, err := db.Query("SELECT * FROM BlogArticle")
	utils.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&blog.ID, &blog.Tag, &blog.Categorie, &blog.Title, &blog.Context, &blog.Author, &blogtime)
		utils.CheckError(err)
		time1, _ := time.ParseInLocation("2006-01-02 15:04:05", blogtime, time.Local)
		blog.Date = LocalTime(time1)
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

//CountBlog 返回记录数
func CountBlog() (count int, err error) {
	db := Getlink()
	defer db.Close()
	err = db.QueryRow("SELECT count(*) FROM BlogArticle").Scan(&count)
	return
}
