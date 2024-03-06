package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transacao struct {
	gorm.Model `json:"-"`
	ClienteID  uint
	Valor      int       `json:"valor"`
	Descricao  string    `json:"descricao"`
	Tipo       string    `json:"tipo"`
	Data       time.Time `json:"data"`
}

type Cliente struct {
	gorm.Model `json:"-"`
	Limite     uint        `json:"limite"`
	Saldo      int         `json:"saldo"`
	Transacoes []Transacao `json:"transacoes"`
}

type ExtratoResposta struct {
	Saldo             Saldo       `json:"saldo"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}

type Saldo struct {
	Total       int       `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      uint      `json:"limite"`
}

type TransacaoResposta struct {
	Limite uint `json:"limite"`
	Saldo  int  `json:"saldo"`
}
