package models
//Response 返回信息结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//JwtToken token令牌
type JwtToken struct{
	Token string `json:"token"`
}