package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateJwt genera el jwt de acceso
func GenerateJwt(issuer string) (string, error) {

	miClave := []byte("MasterDelDesarrollo")

	payload := jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return claims.SignedString(miClave)

}

//ParseJwt manda jwt por cookie
func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("MasterDelDesarrollo"), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
