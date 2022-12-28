package lorry

import "fmt"

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
	fmt.Println("+++ InitBlockchain +++")
	bc := new(blockchain)
	block := bc.CreateBlock(0, "Init hash for genesis block!")
	bc.blockHistory = append(bc.blockHistory, block)

	fmt.Println("--- InitBlockchain ---")
	return *bc
}

func (*blockchain) CreateBlock(nonce int, previousHash string) block {
	fmt.Println("+++ CreateBlock +++")

	t := make([]transaction, 0, 10)
	tx := transaction{
		transationId: "1",
		operations:   operation{},
		nonce:        0,
	}
	t = append(t, tx)
	b := new(block)
	block := b.CreateBlock(t, previousHash)
	block.PrintBlockInfo()

	fmt.Println("--- CreateBlock ---")
	return block
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
	CreateBlock(int, string) block
}

func Blockchain() Blockchainer {
	return &blockchain{
		coinDB:       []account{},
		blockHistory: []block{},
		txDB:         []transaction{},
		faucetCoins:  0,
	}
}
