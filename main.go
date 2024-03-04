package main

import (
	"api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/clientes/:id/transacoes", handlers.PostTransacao)
	r.GET("/clientes/:id/extrato", handlers.GetExtrato)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}