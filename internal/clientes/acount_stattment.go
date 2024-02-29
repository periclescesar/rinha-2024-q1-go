package clientes

import (
	"encoding/json"
	"errors"
)

type Balance struct {
	Total         int    `json:"total,omitempty"`
	Limit         int    `json:"limite,omitempty"`
	StatementDate string `json:"data_extrato,omitempty"`
}

type LastTransaction struct {
	Val         int    `json:"valor,omitempty"`
	Type        string `json:"tipo,omitempty"`
	Description string `json:"descricao,omitempty"`
	RealizedAt  string `json:"realizada_em,omitempty"`
}

type LastTransactions []LastTransaction

func (a *LastTransactions) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type AccountStatement struct {
	Balance          Balance          `json:"saldo,omitempty"`
	LastTransactions LastTransactions `json:"ultimas_transacoes,omitempty"`
}
