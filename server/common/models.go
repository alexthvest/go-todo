package common

import (
	"github.com/alexthvest/go-todo/config"
	"github.com/alexthvest/go-todo/database"
	"github.com/gin-gonic/gin"
)

// Env ...
type Env struct {
	Database *database.Database
	Config   *config.Config
}

// ApiError ...
type ApiError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// MakeErrorReponse ...
func MakeErrorResponse(context *gin.Context, err ApiError) {
	context.JSON(err.StatusCode, err)
	context.Abort()
}
