package ledger

import (
	"github.com/PMoneda/gcoin/block"
	"github.com/PMoneda/gcoin/pow"
	"github.com/PMoneda/gcoin/utils"
)

func (ledger *LedgerBook) AppendBlock(data []byte) (*block.Block, error) {
	head := ledger.GetHeadBlock()
	prev := utils.Int32ToByteArrayNBytes(0, 32)
	if head != nil {
		prev = head.GetHash()
	}
	block := pow.Work(ledger.powDifficult, prev, data)
	err := ledger.attachBlockToLedgerBook(block)
	if err != nil {
		return nil, err
	}
	return block, nil
}
