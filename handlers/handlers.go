package handlers

import (
	"api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTransacao(c *gin.Context) {
	var id int
	if err := GetId(c, &id); err != nil {
		return
	}

	var cliente entities.Cliente
	if err := GetCliente(c, id, &cliente); err != nil {
		return
	}

	var transacao entities.Transacao
	if err := GetTransacaoFromRequest(c, &transacao); err != nil {
		return
	}

	var transacaoResposta entities.TransacaoResposta
	if err := DoTransacao(c, &transacao, &transacaoResposta, &cliente); err != nil {
		return
	}

	c.IndentedJSON(200, transacaoResposta)
}

func GetExtrato(c *gin.Context) {
	c.String(http.StatusOK, "Get extrato")
}
