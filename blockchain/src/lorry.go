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
}
