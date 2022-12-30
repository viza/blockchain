package lorry

import (
	"github.com/gookit/slog"
)

type blockchain struct {
	coinDB       []account
	blockHistory []block
	txDB         []transaction
	faucetCoins  int
}

// GetTokenFromFaucet implements Blockchainer
func (*blockchain) GetTokenFromFaucet(bc *blockchain, acc *account, amount int) {
	slog.Infof("Get some tokens [%d] from faucet", amount)
	acc.UpdateBalance(acc, amount)
	slog.Warn("TODO: Rewrite & clarify logic of faucet tokens")

}

// InitBlockchain implements Blockchainer
func (*blockchain) InitBlockchain() blockchain {
	bc := new(blockchain)
	block := bc.CreateBlock(0, "Init hash for genesis block", nil)
	bc.blockHistory = append(bc.blockHistory, block)
	slog.Warn("TODO: Clarify if genesis block should be with some transactions")

	return *bc
}

func (*blockchain) CreateBlock(nonce int, previousHash string, txArray []transaction) block {
	slog.Info("Create new block")

	b := new(block)
	block := b.CreateBlock(txArray, previousHash)
	//block.PrintBlockInfo()
	return block
}

// ValidateBlock implements Blockchainer
func (*blockchain) ValidateBlock(b block) bool {

	for index, tr := range b.tx {
		slog.Infof("Account [%d] id: %s", index, tr.operations[index].receiver.accountId)
		slog.Infof("Account [%d] balance %d", index, tr.operations[index].sender.balance)
		slog.Infof("Amount of account [%d] to send: %d", index, tr.operations[index].amount)

		if tr.operations[index].receiver.accountId != "" &&
			tr.operations[index].sender.balance > tr.operations[index].amount {
			slog.Warn("Add proper logic and checks in future")
		} else {
			return false
		}
	}
	return true
}

// showCoinDB implements Blockchainer
func (*blockchain) ShowCoinDB(acc []account) {
	var total int = 0
	slog.Info(len(acc))
	for index, value := range acc {
		slog.Info("Account [%d] has balance: %d", index, value.balance)
		total += value.balance
	}
	slog.Info("Total balance of all coins:", total)
}

type Blockchainer interface {
	InitBlockchain() blockchain
	GetTokenFromFaucet(*blockchain, *account, int)
	ValidateBlock(block) bool
	ShowCoinDB([]account)
	CreateBlock(int, string, []transaction) block
	AddBlock2History(*blockchain, block)
	GetAccountBalances(*blockchain) []account
	SetFaucetCoins(*blockchain, int)
	GetFaucetCoins(*blockchain) int
}

func Blockchain() Blockchainer {
	return &blockchain{
		coinDB:       []account{},
		blockHistory: []block{},
		txDB:         []transaction{},
		faucetCoins:  0,
	}
}

func (*blockchain) AddBlock2History(bc *blockchain, b block) {
	bc.blockHistory = append(bc.blockHistory, b)
}

func (*blockchain) GetFaucetCoins(bc *blockchain) int {
	return bc.faucetCoins
}

func (*blockchain) SetFaucetCoins(bc *blockchain, amount int) {
	bc.faucetCoins = amount
}

func (*blockchain) GetAccountBalances(bc *blockchain) []account {
	return bc.coinDB
}
