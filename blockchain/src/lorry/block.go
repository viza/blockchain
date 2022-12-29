package lorry

import (
	"strconv"
	"time"

	"github.com/gookit/slog"
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

	slog.Infof("Block ID:           %s", b.blockId)
	slog.Infof("Previous Hash:      %s", b.prevHash)
	slog.Infof("Time stamp:         %d", b.timeStamp)
	for index, element := range b.tx {
		slog.Infof("Transaction [%d] Id: %s", index, element.transactionId)
	}
}
