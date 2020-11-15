package main

import (
	"github.com/alexthvest/go-todo/common"
	"github.com/alexthvest/go-todo/config"
	"github.com/alexthvest/go-todo/database"
	"github.com/alexthvest/go-todo/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(c.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/todos", func(context *gin.Context) {
		todoRepository := repository.NewTodoRepository(db)
		todos, err := todoRepository.All()

		if err != nil {
			common.MakeErrorResponse(context, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, todos)
	})

	router.POST("/todos", func(context *gin.Context) {
		var todo database.Todo

		todoRepository := repository.NewTodoRepository(db)
		err := context.BindJSON(&todo)

		if err != nil {
			common.MakeErrorResponse(context, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}

		err = todoRepository.Add(todo)
		if err != nil {
			context.JSON(http.StatusBadRequest, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})

			context.Abort()
			return
		}

		context.JSON(http.StatusOK, todo)
	})

	router.Run(c.Port)
}
