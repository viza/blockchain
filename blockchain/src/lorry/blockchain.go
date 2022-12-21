package lorry

type blockchain struct {
	coinDB       []account
	blockHistory []block
	txDB         []transaction
	faucetCoins  int
}

// GetTokenFromFaucet implements Blockchainer
func (*blockchain) GetTokenFromFaucet(acc account, amount int) {
	panic("unimplemented")
}

// InitBlockchain implements Blockchainer
func (*blockchain) InitBlockchain() blockchain {
	panic("unimplemented")
}

// ValidateBlock implements Blockchainer
func (*blockchain) ValidateBlock(b block) {
	panic("unimplemented")
}

// showCoinDB implements Blockchainer
func (*blockchain) showCoinDB() {
	panic("unimplemented")
}

type Blockchainer interface {
	InitBlockchain() blockchain
	GetTokenFromFaucet(acc account, amount int)
	ValidateBlock(b block)
	showCoinDB()
}

func Blockchain() Blockchainer {
	return &blockchain{
		coinDB:       []account{},
		blockHistory: []block{},
		txDB:         []transaction{},
		faucetCoins:  0,
	}
}
