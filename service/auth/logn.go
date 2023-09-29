package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"test-prepare/app"
	"test-prepare/middleware"
	modelDomain "test-prepare/model/domain"
	modelAuth "test-prepare/model/web/auth"
	"test-prepare/repository/utils"
	"time"
)

func LoginService(w http.ResponseWriter, requestLogin modelAuth.LoginRequest) (string, error) {
	var (
		users modelDomain.Users
		err   error
	)

	if err = app.DB.Where("username = ?", requestLogin.Username).First(&users).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := "Username or password wrong"
			utils.Response(w, http.StatusUnauthorized, response, "")
			return "", err
		default:
			utils.Response(w, http.StatusInternalServerError, err.Error(), "")
			return "", err
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(requestLogin.Password)); err != nil {
		response := "Username or password wrong"
		utils.Response(w, http.StatusUnauthorized, response, "")
		return "", err
	}

	expTime := time.Now().Add(time.Minute * 2)
	claims := &middleware.JwtClaim{
		Username: requestLogin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test-prepare",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(middleware.JWT_KEY)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, err.Error(), "")
		return "", err
	}
	//set token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	message := "Login Berhasil"

	return message, nil
}
