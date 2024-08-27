package repositories

import (
	"errors"
	"github.com/periclescesar/rinha-2024-q1-go/internal/clientes"
	"github.com/periclescesar/rinha-2024-q1-go/internal/ports"
	"time"
)

type PostgresClientRepository struct {
}

func (r PostgresClientRepository) GetAccountStatement(id int) (*clientes.AccountStatement, error) {
	db := ports.GetConnection()

	dt := time.Now()

	var accStatement = &clientes.AccountStatement{
		Balance: clientes.Balance{
			StatementDate: dt.Format(time.RFC3339Nano),
		},
	}
	sql := `SELECT
				limite,
				COALESCE(s.valor, 0) as total,
				coalesce(jsonb_agg(t) filter (where t.cliente_id is not null), '[]') as ultimas_transacoes
			FROM clientes c
					 LEFT JOIN saldos s ON c.id = s.cliente_id
					 LEFT JOIN transacoes t ON c.id = t.cliente_id
			WHERE c.id = $1
			GROUP BY c.id, s.valor;`

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)

	err = row.Scan(&accStatement.Balance.Limit, &accStatement.Balance.Total, &accStatement.LastTransactions)

	if err != nil {
		return nil, errors.New("cliente não encontrado")
	}

	return accStatement, nil
}
