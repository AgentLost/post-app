package handler

import "github.com/gin-gonic/gin"

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewErrorMessage(c *gin.Context, statusCode int, message string) {
	response := ErrorMessage{Message: message}
	c.JSON(statusCode, response)
}
