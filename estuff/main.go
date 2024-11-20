package main

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Token from another example.  This token is expired
	var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNTY0MjV9.nVWH9WGX7ejZsgM_fkaVencZvgA6hLXaCgls7M9haMM"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("somerandomsecret"), nil
	})

	switch {
	case token.Valid:
		fmt.Println("Token is valid")
	case errors.Is(err, jwt.ErrTokenMalformed):
		fmt.Println("Token is not correctly set")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		fmt.Println("Signature is invalid")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		fmt.Println("Timing is everything")
	default:
		fmt.Println("Couldn't handle this token:", err)
	}
}
