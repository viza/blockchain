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
	GetPrivateKeyStr(key *ecdsa.PrivateKey) string
	GetPublicKeyStr(key *ecdsa.PublicKey) string
	SignData(key string, data []byte) []byte
	VerifySignature(signature []byte, publicKeyECDSA *ecdsa.PublicKey, data []byte) bool
	GetPublicKey(keys keys) *ecdsa.PublicKey
	GetPrivateKey(keys keys) *ecdsa.PrivateKey
}

func Signature() Signaturer {
	return &keys{
		privateKey: &ecdsa.PrivateKey{},
		publicKey:  &ecdsa.PublicKey{},
	}
}

func (k *keys) genKeyPair() keys {
	fmt.Println("+++ genKeyPair +++")

	prKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
		fmt.Println("--- genKeyPair ---")
	} else {
		k.privateKey = prKey
	}

	publicKey := prKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("  Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	} else {
		k.publicKey = publicKeyECDSA
	}

	fmt.Println("--- genKeyPair ---")
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
	fmt.Println("+++ PrintKeyPair +++")
	privateKeyBytes := crypto.FromECDSA(keys.privateKey)
	fmt.Println("  Private key: " + hexutil.Encode(privateKeyBytes)[2:])

	publicKeyBytes := crypto.FromECDSAPub(keys.publicKey)
	fmt.Println("  Public key: " + hexutil.Encode(publicKeyBytes)[4:])

	fmt.Println("--- PrintKeyPair ---")
}

func (*keys) GetPrivateKeyStr(key *ecdsa.PrivateKey) string {
	return hexutil.Encode(crypto.FromECDSA(key))[2:]
}

func (*keys) GetPublicKeyStr(key *ecdsa.PublicKey) string {
	return hexutil.Encode(crypto.FromECDSAPub(key))[4:]
}

func (*keys) GetPublicKey(keys keys) *ecdsa.PublicKey {
	return keys.publicKey
}

func (*keys) GetPrivateKey(keys keys) *ecdsa.PrivateKey {
	return keys.privateKey
}

// SignData - sign data
func (*keys) SignData(key string, data []byte) []byte {
	fmt.Println("+++ SignData +++")

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

	fmt.Println("  Signed data: " + hexutil.Encode(signature))

	fmt.Println("--- SignData ---")
	return signature
}

func (*keys) VerifySignature(signature []byte, publicKeyECDSA *ecdsa.PublicKey, data []byte) bool {
	fmt.Println("+++ VerifySignature +++")

	publicKey := crypto.FromECDSAPub(publicKeyECDSA)

	hash := crypto.Keccak256Hash(data)
	//fmt.Println("  Hash: ", hash.Hex())

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		fmt.Println("  Error during Ecrecover")
		log.Fatal(err)
		return false
	}

	matches := bytes.Equal(sigPublicKey, publicKey)
	if !matches {
		fmt.Printf("  sigPublicKey (%b) != publicKey (%b)", sigPublicKey, publicKey)
		fmt.Println("  Error during Equal")
		return false
	} else {
		fmt.Println("  Equal: ", matches) // true
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		fmt.Println("  Error during SigToPub")
		log.Fatal(err)
		return false
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKey)
	if !matches {
		fmt.Println("  Error during Match")
		return false
	} else {
		fmt.Println("  Match: ", matches) // true
	}

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKey, hash.Bytes(), signatureNoRecoverID)
	if !verified {
		fmt.Println("  Error during VerifySignature")
		return false
	} else {
		fmt.Println("  VerifySignature: ", verified) // true
	}

	fmt.Println("--- VerifySignature ---")
	return true
}
