package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func SendError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	})
}

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	SendError(c, http.StatusBadRequest, err.Error())
}
