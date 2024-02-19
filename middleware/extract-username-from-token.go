package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

func ExtractUsernameFromToken(tokenString string) (string, error) {
	// Token'ı parse et
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Token'ın imzalı olduğunu kontrol et
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Geçersiz token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	// Token doğrulandı mı?
	if !token.Valid {
		return "", errors.New("Geçersiz token")
	}

	// Kullanıcı adını çıkar
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Token içeriği okunamadı")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("Kullanıcı adı bulunamadı")
	}

	return username, nil
}
