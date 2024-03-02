package entities

type Transacao struct {
	ID        int    `json:"id"`
	Valor     int    `json:"valor"`
	Descricao string `json:"descricao"`
	Tipo      string `json:"tipo"`
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