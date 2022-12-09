package lorry

import (
	"fmt"
	"math/rand"
	"strconv"
)

type account struct {
	accountId string
	wallet    []keys
	balance   int
}

type Accounter interface {
	GenAccount() account
	AddKeyPairToWallet(account, keys)
	GetBalance(account) int
	UpdateBalance(*account, int)
	PrintBalance(account)
}

func Account() Accounter {
	return &account{
		accountId: "",
		wallet:    []keys{},
		balance:   0,
	}
}

// GenAccount - generate pair of keys and create account instance
func (*account) GenAccount() account {
	fmt.Println("+++ GenAccount +++")

	// Generate public and private key
	l := Signature()
	keyPair := l.GenKeyPair()
	keyArr := []keys{keyPair}

	accountId := rand.Intn(100)
	balance := rand.Intn(1000)

	// Print data
	fmt.Printf("  Account id:\n    %d\n", accountId)
	fmt.Printf("  Public key:\n    %s \n", keyArr[0].GetPublicKeyStr(keyPair.publicKey))
	fmt.Printf("  Balance:\n    %d\n", balance)

	fmt.Println("--- GenAccount  ---")
	return account{
		accountId: strconv.Itoa(accountId),
		wallet:    keyArr,
		balance:   balance,
	}
}

// addKeyPairToWallet - add pair of key to account wallet
func (*account) AddKeyPairToWallet(a account, k keys) {
	fmt.Println("+++ AddKeyPairToWallet +++")
	a.wallet = append(a.wallet, k)

	fmt.Printf("  Public key:\n    %s\n", k.GetPublicKeyStr(k.GetPublicKey(k)))
	fmt.Printf("  Private key:\n    %s\n", k.GetPrivateKeyStr(k.GetPrivateKey(k)))

	fmt.Println("--- AddKeyPairToWallet ---")
}

// GetBalance - get account balance
func (*account) GetBalance(xx account) int {
	fmt.Println("+++ GetBalance +++")

	fmt.Println("  Account id:\n   ", xx.accountId)
	fmt.Println("  Balance:\n   ", xx.balance)

	fmt.Println("--- GetBalance ---")
	return xx.balance
}

// UpdateBalance - update balance with new amount
func (*account) UpdateBalance(xx *account, b int) {
	fmt.Println("+++ UpdateBalance +++")

	xx.balance = b
	fmt.Printf("  Balance for account [%s] was updated to [%d]\n", xx.accountId, b)

	fmt.Println("--- UpdateBalance ---")
}

// PrintBalance - print account balance
func (*account) PrintBalance(xx account) {
	fmt.Println("+++ PrintBalance +++")

	fmt.Println("  Account id:\n   ", xx.accountId)
	fmt.Println("  Balance: \n   ", xx.balance)

	fmt.Println("--- PrintBalance ---")
}
