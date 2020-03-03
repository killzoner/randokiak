package controller

import (
	"github.com/gin-gonic/gin"
)

// Controller example
type Controller struct {
}

// IController interface for being able to mock
type IController interface {
	AskMoreProfiles(ctx *gin.Context)
}

// New app
func New() IController {
	return &Controller{}
}
