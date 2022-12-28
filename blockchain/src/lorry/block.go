package lorry

import (
	"fmt"
	"strconv"
	"time"
)

type block struct {
	blockId   string
	prevHash  string
	timeStamp int64
	tx        []transaction
}

// CreateBlock implements Blocker
func (*block) CreateBlock(trx []transaction, prevHash string) block {
	b := new(block)
	b.timeStamp = time.Now().UnixNano()
	b.blockId = strconv.FormatInt(b.timeStamp, 10) // TODO: just for testing purposes
	b.prevHash = prevHash
	b.tx = trx
	return *b
}

type Blocker interface {
	CreateBlock(trx []transaction, prevHash string) block
	PrintBlockInfo()
}

func Block() Blocker {
	return &block{
		blockId:   "",
		prevHash:  "",
		timeStamp: 0,
		tx:        []transaction{},
	}
}

func (b *block) PrintBlockInfo() {
	fmt.Println("+++ PrintBlockInfo +++")

	fmt.Printf("Block ID:       %s\n", b.blockId)
	fmt.Printf("Previous Hash:  %s\n", b.prevHash)
	fmt.Printf("Time stamp:     %d\n", b.timeStamp)
	fmt.Printf("Transactions:   %p\n", b.tx)

	fmt.Println("--- PrintBlockInfo ---")
}
