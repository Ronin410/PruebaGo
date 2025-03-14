package handlers

import "github.com/gin-gonic/gin"

func init() {
	// Inicialización del paquete handlers (si es necesario)
	println("Inicializando el paquete handlers")
}

func Holamundo(c *gin.Context) {
	c.String(200, "Hello")
}

func Holamundo2(c *gin.Context) {
	c.String(200, "Hello handler")
}
