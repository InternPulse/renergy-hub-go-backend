package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/internpulse/renergy-hub-go-backend/config"
	response "github.com/internpulse/renergy-hub-go-backend/pkg"
)

func RequiresAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}
		var jwtSecretKey = os.Getenv("JWT_SECRET")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization header is missing")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Authorization format is invalid")
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		switch {
		case token.Valid:
			claims, ok := token.Claims.(jwt.MapClaims)

			if ok {

				if uID, ok := claims["userID"].(string); ok {
					userID, _ := strconv.Atoi(uID)
					c.Set("user_id", uint(userID))
				} else {
					response.Error(c, http.StatusUnauthorized, "User ID is missing or not a valid number")
					c.Abort()
					return
				}

				if role, ok := claims["role"].(string); ok {
					c.Set("role", role)
				} else {
					response.Error(c, http.StatusUnauthorized, "User role is missing")
					c.Abort()
					return
				}
			} else {
				response.Error(c, http.StatusUnauthorized, "An error occured while parsing the user id")
			}
		case errors.Is(err, jwt.ErrTokenMalformed):
			response.Error(c, http.StatusUnauthorized, "JWT is not correctly set in the Authorization header")
			c.Abort()
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			response.Error(c, http.StatusUnauthorized, "JWT secret is invalid")
			c.Abort()
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			response.Error(c, http.StatusUnauthorized, "JWT has expired")
			c.Abort()
		default:
			response.Error(c, http.StatusUnauthorized, fmt.Sprintf("Couldn't handle this JWT: %v", err))
			c.Abort()
		}

		c.Next()
	}
}
