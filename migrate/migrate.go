package main

import (
	"api/entities"
	"api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&entities.Cliente{})
	initializers.DB.AutoMigrate(&entities.Transacao{})
}
