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

var transactionArray []transaction

type Transactioner interface {
	CreateTransaction([]operation, int) transaction
	GetTransactionArray() []transaction
	PrintTransactionInfo(transaction)
	AddTransaction(transaction)
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

	transactionId := uuid.New() // TODO: shouldb be a hash of agreegated data
	slog.Info("Transation id:   ", transactionId)

	return transaction{transactionId: transactionId.String(), operations: opr, nonce: nonce}
}

func (*transaction) PrintTransactionInfo(t transaction) {
	slog.Info(t.nonce)
	for index, value := range t.operations {
		slog.Infof("Operation [%d] sender account id = %s", index, value.sender.accountId)
		slog.Infof("Operation [%d] sender balance = %d", index, value.sender.balance)
		slog.Infof("Operation [%d] receiver account id = %s", index, value.receiver.accountId)
		slog.Infof("Operation [%d] receiver balance = %d", index, value.receiver.balance)
		slog.Infof("Operation [%d] amount to send = %d", index, value.amount)
		// TODO: maybe add more fields to print later
	}
}

func (*transaction) AddTransaction(t transaction) {
	transactionArray = append(transactionArray, t)
}

func (*transaction) GetTransactionArray() []transaction {
	return transactionArray
}
