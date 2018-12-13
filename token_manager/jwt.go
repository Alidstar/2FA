package tokenmanager

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func generateToken(username, signature string) string {
	claims := &jwt.StandardClaims{
		Issuer:    "cw",
		Audience:  username,
		ExpiresAt: time.Now().Unix() + 2592000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signature))
	if err != nil {
		panic(err)
	}

	return ss
}

func verifyToken(tokenString string, signature string) (bool, string) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signature), nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		fmt.Println("Verify:", claims.Audience, "passed")
		return true, claims.Audience
	} else {
		fmt.Println("Verify: failed")
		return false, ""
	}
}
