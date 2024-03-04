package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTransacao(c *gin.Context) {
	c.String(http.StatusOK, "Post transacao")
}

func GetExtrato(c *gin.Context) {
	c.String(http.StatusOK, "Get extrato")
}