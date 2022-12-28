package lorry

import (
	"math/rand"
	"strconv"

	"github.com/gookit/slog"
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
	PrintAccountInfo(account)
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
	slog.Info("Generate account")

	// Generate public and private key
	l := Signature()
	keyPair := l.GenKeyPair()
	keyArr := []keys{keyPair}

	accountId := rand.Intn(100) //TODO: rewrite with more smart generation
	balance := rand.Intn(1000)  // TODO: generate account with zero balance

	slog.Warn("TODO: implement more smart account id generation")

	// Print data
	slog.Infof("Account id:    %d", accountId)
	slog.Infof("Public key:    %s ", keyArr[0].GetPublicKeyStr(keyPair.publicKey))
	slog.Infof("Balance:       %d", balance)

	return account{
		accountId: strconv.Itoa(accountId),
		wallet:    keyArr,
		balance:   balance,
	}
}

func (*account) PrintAccountInfo(ac account) {
	slog.Info("Account ID:", ac.accountId)
	slog.Info("Balance:", ac.balance)

	for index, element := range ac.wallet {
		slog.Infof("At index %d, wallet key: %d", index, element)
	}
}

// addKeyPairToWallet - add pair of key to account wallet
func (*account) AddKeyPairToWallet(a account, k keys) {
	a.wallet = append(a.wallet, k)

	slog.Infof("Public key:     %s", k.GetPublicKeyStr(k.GetPublicKey(k)))
	slog.Infof("Private key:    %s", k.GetPrivateKeyStr(k.GetPrivateKey(k)))

}

// GetBalance - get account balance
func (*account) GetBalance(xx account) int {

	slog.Info("Account id:    ", xx.accountId)
	slog.Info("Balance:       ", xx.balance)

	return xx.balance
}

// UpdateBalance - update balance with new amount
func (*account) UpdateBalance(xx *account, b int) {

	xx.balance = b
	slog.Infof("Balance for account [%s] was updated to [%d]", xx.accountId, b)

}

// PrintBalance - print account balance
func (*account) PrintBalance(xx account) {

	slog.Info("Account id:     ", xx.accountId)
	slog.Info("Balance:        ", xx.balance)

}
