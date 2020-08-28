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
	SECRET_KEY          = "54f12kkkc8-1fff56-45chg1-a3fsjc-2c546bc2b"
	JWT_ISSUE           = "LINCX_JAVA"
	VERSION_ID = "lincx_v001"
	TokenExpireDuration = time.Hour*24
	PAGE_SIZE = 10
	
)