package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

// CreateToken 创建一个新的 JWT Token
func CreateToken(fooValue string) (string, error) {
	claims := MyCustomClaims{
		fooValue,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("AllYourBase"))
	return signedToken, err
}

func ParseCustomClaims(tokenString string, cl *MyCustomClaims) *jwt.Token {
	token, err := jwt.ParseWithClaims(tokenString, cl, func(token *jwt.Token) (any, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		log.Fatal(err.Error())
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		cl = claims
		fmt.Println(claims.Foo, claims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
	return token
}
