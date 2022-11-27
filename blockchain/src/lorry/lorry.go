package lorry

import (
	"crypto/ecdsa"
	"fmt"

	"log"

	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"golang.org/x/crypto/sha3"
)

type KeyPair struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

var keys = KeyPair{}

func SignData(key string, data []byte) []byte {

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}

	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Signed data: " + hexutil.Encode(signature))
	return signature
}

func PrivateKey2String(privateKey *ecdsa.PrivateKey) string {
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//    fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	return hexutil.Encode(privateKeyBytes)[2:]
}

func GenKeyPair() KeyPair {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	} else {
		keys.PrivateKey = privateKey
	}

	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	} else {
		keys.PublicKey = publicKeyECDSA
	}

	return keys
}

func PrintKeyPair(keys KeyPair) {

	privateKeyBytes := crypto.FromECDSA(keys.PrivateKey)
	fmt.Println("Private key: " + hexutil.Encode(privateKeyBytes)[2:])

	publicKeyBytes := crypto.FromECDSAPub(keys.PublicKey)
	fmt.Println("Public key: " + hexutil.Encode(publicKeyBytes)[4:])
}

func GenAddress(publicKeyECDSA *ecdsa.PublicKey) {
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address: " + address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("Hash = Address: " + hexutil.Encode(hash.Sum(nil)[12:]))
}

func VerifySignature(signature []byte, publicKeyECDSA *ecdsa.PublicKey, data []byte) {

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true

}
