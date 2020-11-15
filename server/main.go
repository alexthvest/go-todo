package main

import (
	"github.com/alexthvest/go-todo/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Run(c.Port)
}
