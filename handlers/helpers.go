package handlers

import (
	"api/entities"
	"api/initializers"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetId(c *gin.Context, id *int) error {
	num, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "O id deve ser um número inteiro")
	}
	*id = num
	return err
}

func GetCliente(c *gin.Context, id int, cliente *entities.Cliente) error {
	result := initializers.DB.First(cliente, id)
	if result.Error != nil {
		c.String(http.StatusNotFound, "Id não encontrado")
	}
	return result.Error
}

func GetTransacaoFromRequest(c *gin.Context, transacao *entities.Transacao) (err error) {
	if err = c.BindJSON(transacao); err != nil {
		c.String(http.StatusUnprocessableEntity, "Corpo da requisição inválido")
	}
	return err
}

func DoTransacao(c *gin.Context, t *entities.Transacao, tr *entities.TransacaoResposta, cl *entities.Cliente) error {
	if len(t.Descricao) < 1 || 10 < len(t.Descricao) {
		c.String(http.StatusUnprocessableEntity, "A descrição deve ter entre 1 e 10 caracteres")
		return errors.New("a descrição deve ter entre 1 e 10 caracteres")
	}

	switch t.Tipo {
	case "c":
		cl.Saldo += t.Valor
		tr.Limite = cl.Limite
		tr.Saldo = cl.Saldo
		cl.Transacoes = append(cl.Transacoes, *t)
		initializers.DB.Save(cl)
		return nil
	case "d":
		if cl.Saldo-t.Valor < -int(cl.Limite) {
			c.String(http.StatusUnprocessableEntity, "O valor da transação ultrapassa o limite")
			return errors.New("o valor da transação ultrapassa o limite")
		}
		cl.Saldo -= t.Valor
		tr.Saldo = cl.Saldo
		tr.Limite = cl.Limite
		cl.Transacoes = append(cl.Transacoes, *t)
		initializers.DB.Save(cl)
		return nil
	default:
		c.String(http.StatusUnprocessableEntity, "O tipo da transação deve ser 'c' ou 'd'")
		return errors.New("o tipo da transação deve ser 'c' ou 'd'")
	}
}
