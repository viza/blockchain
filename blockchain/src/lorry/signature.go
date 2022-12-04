package lorry

import (
	"crypto/ecdsa"
	"fmt"

	"bytes"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type keys struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

type Signaturer interface {
	GenKeyPair() keys
	PrintKeyPair(keys keys)
	GetPrivateKeyStr(keys keys) string
	GetPublicKeyStr(keys keys) string
	SignData(key string, data []byte) []byte
	VerifySignature(signature []byte, publicKeyECDSA *ecdsa.PublicKey, data []byte) bool
	GetPublicKey(keys keys) *ecdsa.PublicKey
}

func Signature() Signaturer {
	return &keys{
		privateKey: &ecdsa.PrivateKey{},
		publicKey:  &ecdsa.PublicKey{},
	}
}

func (k *keys) genKeyPair() keys {

	prKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	} else {
		k.privateKey = prKey
	}

	publicKey := prKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	} else {
		k.publicKey = publicKeyECDSA
	}

	return keys{
		privateKey: k.privateKey,
		publicKey:  publicKeyECDSA,
	}

}

// Generates Key Pair implements Signaturer
func (k *keys) GenKeyPair() keys {
	return k.genKeyPair()
}

// Prints Key Pair implements Signaturer
func (*keys) PrintKeyPair(keys keys) {
	privateKeyBytes := crypto.FromECDSA(keys.privateKey)
	fmt.Println("Private key: " + hexutil.Encode(privateKeyBytes)[2:])

	publicKeyBytes := crypto.FromECDSAPub(keys.publicKey)
	fmt.Println("Public key: " + hexutil.Encode(publicKeyBytes)[4:])
}

func (*keys) GetPrivateKeyStr(keys keys) string {
	return hexutil.Encode(crypto.FromECDSA(keys.privateKey))[2:]
}

func (*keys) GetPublicKeyStr(keys keys) string {
	return hexutil.Encode(crypto.FromECDSAPub(keys.publicKey))[4:]
}

func (*keys) GetPublicKey(keys keys) *ecdsa.PublicKey {
	return keys.publicKey
}

func (*keys) SignData(key string, data []byte) []byte {

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

func (*keys) VerifySignature(signature []byte, publicKeyECDSA *ecdsa.PublicKey, data []byte) bool {

	publicKey := crypto.FromECDSAPub(publicKeyECDSA)

	hash := crypto.Keccak256Hash(data)
	fmt.Println("Hash: ", hash.Hex())

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error during Ecrecover")
		return false
	}

	matches := bytes.Equal(sigPublicKey, publicKey)
	if !matches {
		fmt.Printf("sigPublicKey (%b) != publicKey (%b)", sigPublicKey, publicKey)
		fmt.Println("Error during Equal")
		return false
	} else {
		//fmt.Println(matches) // true
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		fmt.Println("Error during SigToPub")
		log.Fatal(err)
		return false
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKey)
	if !matches {
		fmt.Println("Error during Match")
		return false
	} else {
		//fmt.Println(matches) // true
	}

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKey, hash.Bytes(), signatureNoRecoverID)
	if !verified {
		fmt.Println("Error during VerifySignature")
		return false
	} else {
		//fmt.Println(verified) // true
	}
	return true
}
