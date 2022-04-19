package common

import (
	"gin-go-wp/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//token过期时间 7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	Claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsImV4cCI6MTY1MDk3NzcwMywiaWF0IjoxNjUwMzcyOTAzLCJpc3MiOiJvY2VhbmxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIHRva2VuIn0.smyXxKGhYoRoTdXAqzoyYJkcgPtDBMdMV_p-HrXLkUI
//echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 | base64 -D

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
