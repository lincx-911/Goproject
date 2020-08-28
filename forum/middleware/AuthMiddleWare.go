package middleware

import (
	"fmt"
	"forum/common"

	"github.com/gin-gonic/gin"
)

//AuthMiddleWare 检查权限中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		fmt.Println("进来了")
		tokenstr := ctx.GetHeader("token")
		if tokenstr==""{
			common.HandleAuthError(ctx)
			ctx.Abort()
			return
		}
		claims,err:=common.ParseToken(tokenstr)
		if err!=nil{
			common.HandleAuthError(ctx)
			ctx.Abort()
			return
		}
		ctx.Set("uid",claims.UID)
		ctx.Set("role",claims.Role)
		ctx.Next()
	}
}