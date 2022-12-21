package main

import (
	"blockchain/viza/lorry"
)

func main() {

	// Create sender account
	senderAcc := lorry.Account()
	senderAccount := senderAcc.GenAccount()

	// Update account with 1000
	senderAcc.GetBalance(senderAccount)
	senderAcc.UpdateBalance(&senderAccount, 1000)
	senderAcc.PrintBalance(senderAccount)

	senderAcc.AddKeyPairToWallet(senderAccount, lorry.Signature().GenKeyPair())
	senderAcc.AddKeyPairToWallet(senderAccount, lorry.Signature().GenKeyPair())
	senderAcc.AddKeyPairToWallet(senderAccount, lorry.Signature().GenKeyPair())

	//Create receiver account
	receiverAcc := lorry.Account()
	receiverAccount := receiverAcc.GenAccount()
	//receiverAcc.AddKeyPairToWallet(receiverAccount, receiverKeys)

	//Test data for operation
	sendData := []byte("Some TestData I would like to Send to receiver!!!")

	// Create operation
	operation := lorry.Operation()
	op1 := operation.CreateOperation(senderAccount, receiverAccount, 500, sendData)

	// Verify operation
	op1.VerifyOperation(senderAccount, op1, sendData)

	transaction := lorry.Transaction()

	transaction.CreateTransaction(op1, 150)

	h := lorry.Hash()
	h.ToSHA1("Some data to sha1")

}
