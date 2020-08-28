package common

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//MyClaims 自定义声明
type MyClaims struct{
	UID int `json:"uid"`
	Role string `json:"role"`
	jwt.StandardClaims
}
//var token = "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIwIiwic3ViIjoibGluY3giLCJpc3MiOiJMSU5DWF9KQVZBIiwiZXhwIjoxNTk3OTI0MDY3fQ.C9w3vnDJjfTp8hRy2EM9J2S-eLyUGsM2MW6DGv_dAKk"

//GenToken 生成JWT
func GenToken(id int,name string,role string) (string,error) {
	claim := MyClaims{
		id,
		role,
		jwt.StandardClaims{
			
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    JWT_ISSUE,
			Id:        VERSION_ID,
			Subject:   name,
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return token.SignedString([]byte(SECRET_KEY))
}
//ParseToken 解析JWT
func ParseToken(tokenstr string) (*MyClaims,error) {
	token,err:=jwt.ParseWithClaims(tokenstr,&MyClaims{},func(tokenstring *jwt.Token)(i interface{},err error){
		return []byte(SECRET_KEY),nil
	})
	if err!=nil{
		return nil,err
	}
	if claims,ok:=token.Claims.(*MyClaims);ok&&token.Valid{
		return claims,nil
	}
	return nil,errors.New("invalid token")
}