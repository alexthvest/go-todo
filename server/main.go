package main

import (
	"github.com/alexthvest/go-todo/common"
	"github.com/alexthvest/go-todo/config"
	"github.com/alexthvest/go-todo/database"
	"github.com/alexthvest/go-todo/todos"
	"github.com/gin-gonic/gin"
	"log"
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

	env := &common.Env{Database: db, Config: c}
	router := gin.Default()

	// Todos
	todoGroup := router.Group("/todos")
	todos.RegisterRoutes(env, todoGroup)

	router.Run(c.Port)
}
