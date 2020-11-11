package common

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//MyClaim 自定义声明
type MyClaim struct{
	Appname string `json:"appname"`//app名称
	Appinfo string `json:"appinfo"`//app信息
	Createtime UnixTime `json:"createtime"`//app创建时间
	Roles []int `json:"roles"`//app权限list
	jwt.StandardClaims
}

//TokenExpireDuration token过期时间
const TokenExpireDuration=time.Hour*24

//MySecret 自定义密钥
var MySecret = []byte("54f12kkkc8-1fff56-45chg1-a3fsjc-2c546bc2b")

//GenToken 生成token
func GenToken(claim MyClaim) (string,error) {
	claim.ExpiresAt=time.Now().Add(TokenExpireDuration).Unix()
	claim.IssuedAt=time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

//ParseToken 解析JWT
func ParseToken(tokenstr string) (*MyClaim, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenstr, &MyClaim{}, func(tokenstring *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}