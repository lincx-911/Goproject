package middleswt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webproject/models"

	"github.com/dgrijalva/jwt-go"
)

//ResponseWithJson 响应体转json
func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	fmt.Println("response",w)
}

//GenerateToken 生成token
func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		//"exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
	})
	return token.SignedString([]byte("secret")) //对应的字符串请自行生成，最后足够使用加密后的字符串
}

//TokenMiddleware token中间件
func TokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Cookie")
		fmt.Println(tokenStr)
		if tokenStr == "" {
			ResponseWithJson(w, http.StatusUnauthorized,
				models.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					ResponseWithJson(w, http.StatusUnauthorized,
						models.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
					return nil, fmt.Errorf("not authorization")
				}
				return []byte("secret"), nil
			})
			if !token.Valid {
				ResponseWithJson(w, http.StatusUnauthorized,
					models.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
			} else {
				next(w, r)
			}
		}
	}
}
