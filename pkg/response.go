package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, message string, data ...interface{}) {
	response := Response{
		Status:  "success",
		Message: message,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	c.JSON(code, response)
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Status:  "error",
		Message: message,
	})
}
