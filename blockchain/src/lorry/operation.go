package lorry

import (
	"fmt"
)

type operation struct {
	sender    account
	receiver  account
	amount    int
	signature []byte
}

type Operationer interface {
	CreateOperation(account, account, []byte, int) operation
	VerifyOperation(account, operation) bool
}

func Operation() Operationer {
	return &operation{
		sender:    account{},
		receiver:  account{},
		amount:    0,
		signature: []byte{},
	}
}

func (*operation) VerifyOperation(snd account, op operation) bool {
	fmt.Println("=== VerifyOperation ===")

	fmt.Println(op.amount)
	balance := snd.GetBalance(snd)
	fmt.Println("Balance of sender:", balance)
	//check amount
	if balance >= op.amount {
		fmt.Printf("Will be sent %d from sender with balance %d\n", op.amount, balance)
	} else {
		fmt.Println("Not enough money on balance")
		return false
	}

	//TODO: check signature
	return true
}

func (*operation) CreateOperation(sender account, receiver account, data []byte, amount int) operation {
	fmt.Println("=== CreateOperation ===")

	fmt.Println("Amount to send: ", amount)

	return operation{
		sender:    sender,
		receiver:  receiver,
		amount:    amount,
		signature: data,
	}
}
