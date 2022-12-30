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
func (*blockchain) GetTokenFromFaucet(acc account, amount int) {
	slog.Error("TODO: implement")
}

// InitBlockchain implements Blockchainer
func (*blockchain) InitBlockchain() blockchain {
	bc := new(blockchain)
	block := bc.CreateBlock(0, "Init hash for genesis block!", nil)
	bc.blockHistory = append(bc.blockHistory, block)

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
func (*blockchain) ValidateBlock(b block) {
	slog.Error("TODO: implement")

}

// showCoinDB implements Blockchainer
func (*blockchain) showCoinDB() {
	slog.Error("TODO: implement")
}

type Blockchainer interface {
	InitBlockchain() blockchain
	GetTokenFromFaucet(acc account, amount int)
	ValidateBlock(b block)
	showCoinDB()
	CreateBlock(int, string, []transaction) block
	AddBlock2History(blockchain, block)
}

func Blockchain() Blockchainer {
	return &blockchain{
		coinDB:       []account{},
		blockHistory: []block{},
		txDB:         []transaction{},
		faucetCoins:  0,
	}
}

func (*blockchain) AddBlock2History(bc blockchain, b block) {
	bc.blockHistory = append(bc.blockHistory, b)
}
