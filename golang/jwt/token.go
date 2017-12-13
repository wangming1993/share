package main

//学习 github.com/dgrijalva/jwt-go

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//signKey 必须为 []byte 类型， 否则会报 key is invalid 错误
var signKey = []byte("mike")

func main() {
	tokenString := GenerateToken()

	ParseTokenString(tokenString)
}

//GenerateToken generate jwt token
func GenerateToken() string {
	now := time.Now().Unix()
	claims := jwt.MapClaims{
		"appId": "bar",
		"exp":   now + int64(60*60),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signKey)

	panicIfError(err)

	fmt.Println("Generate jwt token --->>>")
	fmt.Println(tokenString)

	return tokenString
}

//ParseTokenString parse a given token string
func ParseTokenString(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return signKey, nil
	})

	panicIfError(err)

	if err = token.Claims.Valid(); err != nil {
		fmt.Println("claim not vaild")
	}

	fmt.Println("Parse jwt token --->>>")
	fmt.Println(token.Claims, token.Header)
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
