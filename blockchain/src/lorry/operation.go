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
	CreateOperation(account, account, int, []byte) operation
	VerifyOperation(account, operation, []byte) bool
}

func Operation() Operationer {
	return &operation{
		sender:    account{},
		receiver:  account{},
		amount:    0,
		signature: []byte{},
	}
}

// CreateOperation - create operation
func (*operation) CreateOperation(sender account, receiver account, amount int, data []byte) operation {
	fmt.Println("+++ CreateOperation +++")

	fmt.Println("  Amount to send: ", amount)

	//Sign the data
	//TODO: clarify if many private keys, which one is used for signature
	signed := Signature().SignData(sender.wallet[0].GetPrivateKeyStr(sender.wallet[0].privateKey), data)

	//sender.wallet.GetPublicKeyStr(sender.wallet[0].GetPublicKey(sender.wallet[0]))

	fmt.Println("--- CreateOperation ---")
	return operation{
		sender:    sender,
		receiver:  receiver,
		amount:    amount,
		signature: signed,
	}
}

// VerifyOperation - verify validity of operation (available balance and signature)
func (*operation) VerifyOperation(snd account, op operation, signedData []byte) bool {
	fmt.Println("+++ VerifyOperation +++")

	balance := snd.GetBalance(snd)
	//fmt.Println("  Balance of sender:", balance)

	//check amount
	if balance >= op.amount {
		fmt.Printf("  Will be sent %d from sender with balance %d\n", op.amount, balance)
	} else {
		fmt.Println("  Not enough money on balance")
		fmt.Println("--- VerifyOperation ---")
		return false
	}

	for i := range snd.wallet {
		Signature().VerifySignature(op.signature, snd.wallet[i].GetPublicKey(snd.wallet[i]), signedData)
	}

	fmt.Println("--- VerifyOperation ---")
	return true
}
