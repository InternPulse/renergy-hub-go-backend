package utils

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetUserID(c *gin.Context) (uint, error) {
	userId, exists := c.Get("user_id")
	if !exists {
		return 0, fmt.Errorf("user id not found in request")
	}

	switch id := userId.(type) {
	case uint:
		return id, nil
	case int64:
		return uint(id), nil
	default:
		return 0, fmt.Errorf("invalid user ID type")
	}
}

func GetGenericId(c *gin.Context, param string) (uint, error) {
	paramStr := c.Param(param)
	genericId := uint(0)
	_, err := fmt.Sscanf(paramStr, "%d", &genericId)
	if err != nil {
		return 0, err
	}
	return genericId, nil
}
