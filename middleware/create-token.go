package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(username string) (string, error) {
	// Token'ı oluştur
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 saat geçerli olan bir token

	// Token'ı imzala
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
