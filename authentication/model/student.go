package model

//Student 学生
type Student struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Score float64 `json:"score"`
}
//Course 课程
type Course struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Teacher string `json:"teacher"`
}
//Teacher 课程
type Teacher struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Department string `json:"department"`
}
