package main

import (
	"log"

	"PruebaGo/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/holamundo", handlers.Holamundo)
	r.GET("/holamundo2", handlers.Holamundo2)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
