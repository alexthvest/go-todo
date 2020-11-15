package main

import (
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
			log.Fatal(err)

			context.JSON(http.StatusBadRequest, err)
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, todos)
	})

	router.Run(c.Port)
}
