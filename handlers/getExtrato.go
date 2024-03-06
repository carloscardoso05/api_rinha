package handlers

import (
	"api/entities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetExtrato(c *gin.Context) {
	var id int
	if err := GetId(c, &id); err != nil {
		return
	}

	var cliente entities.Cliente
	if err := GetCliente(c, id, &cliente); err != nil {
		return
	}

	var ultimas []entities.Transacao
	if err := GetUltimasTransacoes(c, id, &ultimas); err != nil {
		return
	}

	extrato := entities.ExtratoResposta{
		Saldo: entities.Saldo{
			Total:       cliente.Saldo,
			DataExtrato: time.Now(),
			Limite:      cliente.Limite,
		},
		UltimasTransacoes: ultimas,
	}

	c.IndentedJSON(http.StatusOK, extrato)
}
