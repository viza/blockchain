package lorry

import (
	"github.com/google/uuid"
	"github.com/gookit/slog"
)

type transaction struct {
	transactionId string
	operations    []operation
	nonce         int
}

type Transactioner interface {
	CreateTransaction([]operation, int) transaction
}

func Transaction() Transactioner {
	return &transaction{
		transactionId: "",
		operations:    []operation{},
		nonce:         0,
	}
}

// CreateTransaction - create transation
func (*transaction) CreateTransaction(opr []operation, nonce int) transaction {

	transactionId := uuid.New()
	slog.Info("Transation id:   ", transactionId)

	return transaction{transactionId: transactionId.String(), operations: opr, nonce: nonce}
}
