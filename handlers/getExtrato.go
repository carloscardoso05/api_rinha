package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExtrato(c *gin.Context) {
	c.String(http.StatusOK, "Get extrato")
}