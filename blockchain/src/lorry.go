package main

import (
	"blockchain/viza/lorry"
	"fmt"
)

func main() {
	fmt.Println("Lorry blockchain")

	// with constructor
	l := lorry.Signature()
	keys := l.GenKeyPair()
	l.PrintKeyPair(keys)

	privateKey := l.GetPrivateKeyStr(keys)

	data := []byte("Some TestData I would like to Sign")

	// Generate signature
	signature := l.SignData(privateKey, data)

	//publicKey := l.GetPublicKeyStr(keys)
	publicKey := l.GetPublicKey(keys)

	//Verify signature
	res := l.VerifySignature(signature, publicKey, data)
	if !res {
		fmt.Println("Error: Wrong Signature! It doesn't pass verififcation")
	} else {
		fmt.Println("Sinature has been successfully verified!")
	}

}
