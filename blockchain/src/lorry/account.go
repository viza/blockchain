package lorry

type account struct {
	accountId string
	wallet    []keys
	balance   int
}

// addKeyPairToWallet implements Accounter
func (*account) AddKeyPairToWallet(keys) {
	panic("unimplemented")
}

// genAccount implements Accounter
func (*account) GenAccount() account {
	l := Signature()
	keyPair := l.GenKeyPair()
	keyArr := []keys{keyPair}

	return account{
		accountId: "",
		wallet:    keyArr,
		balance:   0,
	}
}

// getBalance implements Accounter
func (*account) GetBalance() {
	panic("unimplemented")
}

// printBalance implements Accounter
func (*account) PrintBalance() {
	panic("unimplemented")
}

// updateBalance implements Accounter
func (*account) UpdateBalance(int) {
	panic("unimplemented")
}

type Accounter interface {
	GenAccount() account
	AddKeyPairToWallet(keys)
	UpdateBalance(int)
	GetBalance()
	PrintBalance()
}

func Account() Accounter {
	return &account{
		accountId: "",
		wallet:    []keys{},
		balance:   0,
	}
}
