package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewErrorMessageResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorMessage{Message: message})
}
