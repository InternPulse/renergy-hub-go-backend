package main

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNTUwNTZ9.PLITo7BsaQ9rFwpBscu9tgYhxZC1p20yGWRX-yPv7v0"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("somerandomsecret"), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["exp"], claims["user_id"])
	} else {
		fmt.Println(err)
	}
}
