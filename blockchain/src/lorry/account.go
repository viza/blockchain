package lorry

import "fmt"

type account struct {
	accountId string
	wallet    []keys
	balance   int
}

// addKeyPairToWallet implements Accounter
func (*account) AddKeyPairToWallet(a account, k keys) {
	fmt.Println("=== AddKeyPairToWallet ===")
	a.wallet = append(a.wallet, k)
	/* for i := range a.wallet {
		fmt.Printf("Key %d = %s", i, a.wallet[i].publicKey)
	} */
}

// genAccount implements Accounter
func (*account) GenAccount() account {
	fmt.Println("=== GenAccount ===")
	l := Signature()
	keyPair := l.GenKeyPair()
	keyArr := []keys{keyPair}
	//fmt.Println(keyArr[0].privateKey)

	return account{
		accountId: "12345", //just for test, TODO: Generate unique ID
		wallet:    keyArr,
		balance:   100, // just for test, TODO: set up 0
	}
}

// getBalance implements Accounter
func (*account) GetBalance(xx account) int {
	fmt.Println("=== GetBalance ===")
	//fmt.Println("Account id: ", xx.accountId)
	//fmt.Println("Balance: ", xx.balance)
	return xx.balance
}

// printBalance implements Accounter
func (*account) PrintBalance(xx account) {
	fmt.Println("=== PrintBalance ===")

	fmt.Println("Account balance...")

	fmt.Println("Account id: ", xx.accountId)
	fmt.Println("Balance: ", xx.balance)
}

// updateBalance implements Accounter
func (*account) UpdateBalance(xx *account, b int) {
	fmt.Println("=== UpdateBalance ===")

	xx.balance = b
}

type Accounter interface {
	GenAccount() account
	AddKeyPairToWallet(account, keys)
	UpdateBalance(*account, int)
	GetBalance(account) int
	PrintBalance(account)
}

func Account() Accounter {
	return &account{
		accountId: "",
		wallet:    []keys{},
		balance:   0,
	}
}
