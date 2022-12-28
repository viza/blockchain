package main

import (
	"blockchain/viza/lorry"

	"github.com/gookit/slog"
)

func main() {

	slog.Info("Main function with test code of blockchain...")

	// Create sender account
	slog.Info("Create sender account")
	senderAcc := lorry.Account()
	senderAccount := senderAcc.GenAccount()
	senderAcc.PrintAccountInfo(senderAccount)
	senderAcc.PrintBalance(senderAccount)

	// Update account with 1000
	slog.Info("Update sender's balance with 1000")
	senderAcc.GetBalance(senderAccount)
	senderAcc.UpdateBalance(&senderAccount, 1000)
	senderAcc.PrintBalance(senderAccount)

	slog.Info("Generate and add keys to sender's wallet")
	senderAcc.AddKeyPairToWallet(senderAccount, lorry.Signature().GenKeyPair())
	senderAcc.AddKeyPairToWallet(senderAccount, lorry.Signature().GenKeyPair())
	senderAcc.AddKeyPairToWallet(senderAccount, lorry.Signature().GenKeyPair())

	//Create receiver account
	slog.Info("Create receiver account")
	receiverAcc := lorry.Account()
	receiverAccount := receiverAcc.GenAccount()
	//receiverAcc.AddKeyPairToWallet(receiverAccount, receiverKeys)

	//Test data for operation
	sendData := []byte("Some TestData I would like to Send to receiver!!!")

	// Create operation
	slog.Info("Create operation")
	operation := lorry.Operation()
	op1 := operation.CreateOperation(senderAccount, receiverAccount, 500, sendData)
	//op2 := operation.CreateOperation(senderAccount, receiverAccount, 800, sendData)

	// Verify operation
	slog.Info("Verify operation")
	op1.VerifyOperation(senderAccount, op1, sendData)

	//	t := lorry.Transaction()
	//	fmt.Println("Type:", reflects.TypeOf(t))

	//	t := lorry.Transaction()
	//	tx := t.CreateTransaction(op1, 150)
	//	tx2 := t.CreateTransaction(op2, 200)

	//	h := lorry.Hash()
	//	HashStr := h.ToSHA1("Some data to sha1")

	//b := lorry.Block()
	//b.CreateBlock(reflects.Type{tx, tx2}, HashStr)

	// Create blockchain
	slog.Info("Create blockchain")
	blockchain := lorry.Blockchain()
	blockchain.InitBlockchain()
	b0 := blockchain.CreateBlock(0, "InitHash")
	blockchain.ValidateBlock(b0)

}
