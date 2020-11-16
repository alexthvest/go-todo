package todos

import (
	"github.com/alexthvest/go-todo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterRoutes ...
func RegisterRoutes(env *common.Env, router *gin.RouterGroup) {
	router.GET("/", getTodos(env))
	router.POST("/", addTodo(env))
}

// getTodos ...
func getTodos(env *common.Env) gin.HandlerFunc {
	return func(context *gin.Context) {
		todoRepository := NewRepository(env.Database)
		ts, err := todoRepository.All()

		if err != nil {
			common.MakeErrorResponse(context, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, ts)
	}
}

// addTodo ...
func addTodo(env *common.Env) gin.HandlerFunc {
	return func(context *gin.Context) {
		var todo Todo

		todoRepository := NewRepository(env.Database)
		err := context.BindJSON(&todo)

		if err != nil {
			common.MakeErrorResponse(context, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}

		doc, err := todoRepository.Add(todo)
		if err != nil {
			common.MakeErrorResponse(context, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, doc)
	}
}
