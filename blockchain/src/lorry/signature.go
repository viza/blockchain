package lorry

import (
	"crypto/ecdsa"

	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gookit/slog"
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

	prKey, err := crypto.GenerateKey()
	if err != nil {
		slog.Fatal(err)
	} else {
		k.privateKey = prKey
	}

	publicKey := prKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		slog.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
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
	slog.Info("Private key:   " + hexutil.Encode(privateKeyBytes)[2:])

	publicKeyBytes := crypto.FromECDSAPub(keys.publicKey)
	slog.Info("Public key:    " + hexutil.Encode(publicKeyBytes)[4:])

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

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		slog.Fatal("Error during hex to ecdsa", err)
	}

	hash := crypto.Keccak256Hash(data)
	slog.Info("Hash:   ", hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		slog.Fatal("Error during sign", err)
	}

	slog.Info("Signed data:   " + hexutil.Encode(signature))

	return signature
}

func (*keys) VerifySignature(signature []byte, publicKeyECDSA *ecdsa.PublicKey, data []byte) bool {

	publicKey := crypto.FromECDSAPub(publicKeyECDSA)

	hash := crypto.Keccak256Hash(data)
	slog.Info("Hash:   ", hash.Hex())

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		slog.Fatal("Error during Ecrecover", err)
		return false
	}

	matches := bytes.Equal(sigPublicKey, publicKey)
	if !matches {
		slog.Errorf("sigPublicKey (%b) != publicKey (%b)", sigPublicKey, publicKey)
		return false
	} else {
		slog.Info("Equal:   ", matches) // true
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		slog.Error("Error during SigToPub", err)
		return false
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKey)
	if !matches {
		slog.Error("Error during Match", matches)
		return false
	} else {
		slog.Info("Match:    ", matches) // true
	}

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKey, hash.Bytes(), signatureNoRecoverID)
	if !verified {
		slog.Error("Error during VerifySignature", verified)
		return false
	} else {
		slog.Info("VerifySignature: ", verified) // true
	}

	return true
}
