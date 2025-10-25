package main

import (
	"log"

	"peddie-backend/functions/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	models.ConnectDatabase()

	log.Println("(___ Server Has Started ___)")
	r.Run()
}
