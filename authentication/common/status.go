package common

const(
	//OK 请求成功
	OK              = iota
	ParamsError    // 传参错误
	ServerError    // 服务错误
	OperationError // 操作错误
	RecordExists   // 记录已存在
	RecordNotFound // 记录不存在
	AuthError      // 需要登陆
)
//Result 请求响应体 
type Result struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}