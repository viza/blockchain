package lorry

import (
	"math/rand"
	"strconv"

	"github.com/gookit/slog"
)

type transaction struct {
	transationId string
	operations   []operation
	nonce        int
}

type Transactioner interface {
	CreateTransaction([]operation, int) transaction
}

func Transaction() Transactioner {
	return &transaction{
		transationId: "",
		operations:   []operation{},
		nonce:        0,
	}
}

// CreateTransaction - create transation
func (*transaction) CreateTransaction(opr []operation, nonce int) transaction {

	transactionId := strconv.Itoa(rand.Intn(100))
	slog.Info("Transation id:   ", transactionId)

	return transaction{transationId: transactionId, operations: opr, nonce: nonce}
}
