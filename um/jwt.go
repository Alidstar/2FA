package um

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func (um *UserManager) generateToken(username string) string {
	claims := &jwt.StandardClaims{
		Issuer:   "cw",
		Audience: username,
		Subject:  "cwauthen",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(um.signature))
	if err != nil {
		panic(err)
	}

	return ss
}

func (um *UserManager) verifyToken(tokenString string) (bool, string) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(um.signature), nil
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
