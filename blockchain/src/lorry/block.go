package lorry

import "fmt"

type block struct {
	blockId  string
	prevHash string
	trx      transaction
}

// CreateBlock implements Blocker
func (*block) CreateBlock(trx transaction, prevHash string) block {
	var b block
	fmt.Println("TODO: Create block with all necessary elements")
	b.prevHash = prevHash
	b.blockId = "0001"
	b.trx = trx

	return b
}

type Blocker interface {
	CreateBlock(trx transaction, prevHash string) block
}

func Block() Blocker {
	return &block{
		blockId:  "",
		prevHash: "",
		trx:      transaction{},
	}
}
