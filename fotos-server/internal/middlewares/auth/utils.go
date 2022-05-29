package auth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id uint
	jwt.StandardClaims
}

func GenerateToken(id uint) string {
	var payload Claims
	currTime := time.Now()

	payload.Id = id
	payload.IssuedAt = currTime.Unix()
	payload.ExpiresAt = time.Now().Add(time.Minute * 120).Unix()

	secret := []byte("yourmom")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, error := token.SignedString(secret)

	if error != nil {
		fmt.Println(error.Error())
	}

	return tokenString
}

func VerifyToken(tokenString string) (bool, uint) {
	payload := Claims{}
	secret := []byte("yourmom")

	token, err := jwt.ParseWithClaims(tokenString, &payload, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return false, 0
	}

	return token.Valid, payload.Id
}
