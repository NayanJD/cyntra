package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type InvalidJwt struct{}

func (e *InvalidJwt) Error() string {
	return "The jwt is invalid"
}

func CreateToken(claims jwt.Claims, d time.Duration, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secret string, claim jwt.Claims) (*jwt.Token, error) {
	token, _ := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		fmt.Println("secret: ", secret)

		return []byte(secret), nil
	})

	if token.Valid {
		return token, nil
	} else {
		fmt.Println("token.Valid: ", token.Valid)
		return nil, &InvalidJwt{}
	}
}
