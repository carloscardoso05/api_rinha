package entities

import "time"

type Transacao struct {
	ID        int       `json:"id"`
	Valor     int       `json:"valor"`
	Descricao string    `json:"descricao"`
	Tipo      string    `json:"tipo"`
	Data      time.Time `json:"data"`
}

type Extrato struct {
	Saldo             Saldo       `json:"saldo"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}

type Saldo struct {
	Total       int       `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      int       `json:"limite"`
}

type Cliente struct {
	ID         int         `json:"id"`
	Limite     int         `json:"limite"`
	Saldo      int         `json:"saldo"`
	Transacoes []Transacao `json:"transacoes"`
}

type TransacaoResposta struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}
