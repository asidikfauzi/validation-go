package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"test-prepare/app"
	"test-prepare/repository/utils"
)

var JWT_KEY = []byte(app.GetEnv("JWT_KEY"))

type JwtClaim struct {
	Username string
	jwt.RegisteredClaims
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			fmt.Println(err)
			if err == http.ErrNoCookie {
				response := "Unauthorized"
				utils.Response(w, http.StatusUnauthorized, response, "")
				return
			}
		}
		tokenString := c.Value

		claims := &JwtClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := "Unauthorized"
				utils.Response(w, http.StatusUnauthorized, response, "")
				return
			case jwt.ValidationErrorExpired:
				response := "Unauthorized, Token Expired"
				utils.Response(w, http.StatusUnauthorized, response, "")
				return
			default:
				response := "Unauthorized"
				utils.Response(w, http.StatusUnauthorized, response, "")
				return
			}
		}
		if !token.Valid {
			response := "Unauthorized, Token Expired"
			utils.Response(w, http.StatusUnauthorized, response, "")
			return
		}

		next.ServeHTTP(w, r)
	})
}
