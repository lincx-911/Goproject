package middleware

import (
	"authentication/common"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//AuthMiddleware 验证中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		token := cookie.Value
		claim, err := common.ParseToken(token)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// requestBody:=new(bytes.Buffer)
		// json.NewEncoder(requestBody).Encode(claim)
		body, err := json.Marshal(claim.Roles)
		r.Body = ioutil.NopCloser(strings.NewReader(string(body)))
		next.ServeHTTP(rw, r)
	})
}
