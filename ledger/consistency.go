package ledger

import (
	"bytes"
	"fmt"

	"github.com/PMoneda/bblock/block"
	"github.com/PMoneda/bblock/utils"
)

/*
Este metodo faz a validacao de todo o arquivo de blockchain para validar se algum bloco foi violado

*/
func (ledger *LedgerBook) CheckConsistency() error {
	head := ledger.GetHeadBlock()
	next := head.GetPreviousHash()
	for bytes.Compare(head.GetHash(), utils.Int32ToByteArrayNBytes(0, 32)) != 0 {
		//fmt.Printf("Current block hash:%x\nCurrent block Previous hash:%x\nCurrent block body: %s\n\n", head.GetHash(), head.GetPreviousHash(), head.GetBody())
		head = ledger.GetBlock(head.GetPreviousHash())
		if head == nil {
			break
		}
		if bytes.Compare(head.GetHash(), next) != 0 || !block.CheckBlockConsistency(head) {
			return fmt.Errorf("block %x is corrupted", head.GetHash())
		}
		next = head.GetPreviousHash()
	}

	return nil
}
