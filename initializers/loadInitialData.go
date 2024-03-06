package initializers

import (
	"api/entities"
	"log"

	"gorm.io/gorm"
)

func LoadInitialData() {
	DB.Exec("DELETE FROM transacaos")
	DB.Exec("DELETE FROM clientes")

	for _, c := range clientes {
		result := DB.Create(&c)
		if result.Error != nil {
			log.Fatal("Erro ao inicializar os dados")
		}
	}
}

var clientes []entities.Cliente = []entities.Cliente{
	{
		Model: gorm.Model{
			ID: 1,
		},
		Limite:     100000,
		Saldo:      0,
		Transacoes: []entities.Transacao{},
	},
	{
		Model: gorm.Model{
			ID: 2,
		},
		Limite:     80000,
		Saldo:      0,
		Transacoes: []entities.Transacao{},
	},
	{
		Model: gorm.Model{
			ID: 3,
		},
		Limite:     1000000,
		Saldo:      0,
		Transacoes: []entities.Transacao{},
	},
	{
		Model: gorm.Model{
			ID: 4,
		},
		Limite:     10000000,
		Saldo:      0,
		Transacoes: []entities.Transacao{},
	},
	{
		Model: gorm.Model{
			ID: 5,
		},
		Limite:     500000,
		Saldo:      0,
		Transacoes: []entities.Transacao{},
	},
}
