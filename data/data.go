package data
import "api/entities"

// dados inciais
var Clientes []entities.Cliente = []entities.Cliente{
	{
		ID:     1,
		Limite: 100000,
		Saldo:  0,
		Transacoes: []entities.Transacao{},
	},
	{
		ID:     2,
		Limite: 80000,
		Saldo:  0,
		Transacoes: []entities.Transacao{},
	},
	{
		ID:     3,
		Limite: 1000000,
		Saldo:  0,
		Transacoes: []entities.Transacao{},
	},
	{
		ID:     4,
		Limite: 10000000,
		Saldo:  0,
		Transacoes: []entities.Transacao{},
	},
	{
		ID:     5,
		Limite: 500000,
		Saldo:  0,
		Transacoes: []entities.Transacao{},
	},
}