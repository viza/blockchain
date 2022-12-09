package lorry

import (
	"fmt"
	"math/rand"
	"strconv"
)

type transaction struct {
	transationId string
	operations   operation
}

type Transactioner interface {
	CreateTransaction(operation, int) transaction
}

func Transaction() Transactioner {
	return &transaction{
		transationId: "",
		operations:   operation{},
	}
}

// CreateTransaction - create transation
func (*transaction) CreateTransaction(opr operation, nonce int) transaction {
	fmt.Println("+++ CreateTransaction +++")

	transactionId := strconv.Itoa(rand.Intn(100))
	fmt.Println("  Transation id:\n   ", transactionId)

	fmt.Println("--- CreateTransaction ---")
	return transaction{transationId: transactionId, operations: opr}
}
