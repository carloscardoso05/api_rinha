package main

import (
	"api/handlers"
	"api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.LoadInitialData()
}

func main() {
	r := gin.Default()
	r.POST("/clientes/:id/transacoes", handlers.PostTransacao)
	r.GET("/clientes/:id/extrato", handlers.GetExtrato)
	r.Run()
}
