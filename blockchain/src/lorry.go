package main

import (
	"blockchain/viza/lorry"
	"fmt"
)

func main() {
	fmt.Println("Lorry blockchain")

	// Generate keys
	l := lorry.Signature()
	keys := l.GenKeyPair()
	l.PrintKeyPair(keys)

	privateKey := l.GetPrivateKeyStr(keys)

	data := []byte("Some TestData I would like to Sign!!!")

	// Generate signature
	signature := l.SignData(privateKey, data)

	publicKey := l.GetPublicKey(keys)

	//Verify signature
	res := l.VerifySignature(signature, publicKey, data)
	if !res {
		fmt.Println("Error: Wrong Signature! It doesn't pass verififcation")
	} else {
		fmt.Println("Sinature has been successfully verified!")
	}

	// Create account
	//fmt.Println("Create account...")
	a := lorry.Account()
	account := a.GenAccount()
	//	fmt.Println("Account created...")

	// Get balance
	balance := a.GetBalance(account)
	fmt.Println("Balance: ", balance)

	// Add key pair to account wallet
	a.AddKeyPairToWallet(account, keys)

	// Generate second pair of keys for testing
	keys2 := l.GenKeyPair()
	// Add second key pair to account wallet
	a.AddKeyPairToWallet(account, keys2)

	// Print balance
	a.PrintBalance(account)

	//Update balance
	a.UpdateBalance(&account, 250)
	a.PrintBalance(account)

	//Create operation

	//Create sender account
	senderSignature := lorry.Signature()
	senderKeys := senderSignature.GenKeyPair()
	senderAcc := lorry.Account()
	senderAccount := senderAcc.GenAccount()
	senderAcc.AddKeyPairToWallet(senderAccount, senderKeys)

	//Create receiver account
	receiverSignature := lorry.Signature()
	receiverKeys := receiverSignature.GenKeyPair()
	receiverAcc := lorry.Account()
	receiverAccount := receiverAcc.GenAccount()
	receiverAcc.AddKeyPairToWallet(receiverAccount, receiverKeys)

	//Generate signature by sender
	sendData := []byte("Some TestData I would like to Send to receiver!!!")

	// Generate signature
	signedData := senderSignature.SignData(senderKeys.GetPrivateKeyStr(senderKeys), sendData)

	operation := lorry.Operation()
	op1 := operation.CreateOperation(senderAccount, receiverAccount, signedData, 500)
	op1.VerifyOperation(senderAccount, op1)
}
