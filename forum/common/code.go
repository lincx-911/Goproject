package common

import "time"

const (
	OK             = iota
	ParamsError    // 传参错误
	ServerError    // 服务错误
	OperationError // 操作错误
	RecordExists   // 记录已存在
	RecordNotFound // 记录不存在
	AuthError      // 需要登陆
)

const (
	HEADER_TOKEN        = "token"
	HEADER_VERSION      = "version"
	SECRET_KEY          = "54f12kkkc8-1fxxxxxxxxxxxxxxxxxx"
	JWT_ISSUE           = "LINCXxxxxxx"
	VERSION_ID = "lincxxxxxxxxx1"
	TokenExpireDuration = time.Hour*24
	PAGE_SIZE = 10
	
)
