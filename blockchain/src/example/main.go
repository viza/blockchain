package main

import (
	"blockchain/lorry"
	"fmt"
)

func main() {
	// Generate keys
	keys := lorry.GenKeyPair()
	lorry.PrintKeyPair(keys)

	//lorry.GenAddress(keys.PublicKey)

	data := []byte("Some TestData 2 Sign")
	privateKeyStr := lorry.PrivateKey2String(keys.PrivateKey)
	fmt.Println("Private key string: " + privateKeyStr)

	// Generate signature
	signature := lorry.SignData(privateKeyStr, data)

	//Verify signature
	lorry.VerifySignature(signature, keys.PublicKey, data)
}
