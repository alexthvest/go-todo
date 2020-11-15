package common

import (
	"github.com/gin-gonic/gin"
)

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
