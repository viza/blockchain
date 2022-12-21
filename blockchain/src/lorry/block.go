package lorry

import (
	"strconv"
	"time"
)

type block struct {
	blockId  string
	prevHash string
	trx      transaction
}

// CreateBlock implements Blocker
func (*block) CreateBlock(trx transaction, prevHash string) block {
	b := new(block)
	t := time.Now().UnixNano()
	b.blockId = strconv.FormatInt(t, 10) // TODO: just for testing purposes
	b.prevHash = prevHash
	b.trx = trx
	return *b
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
