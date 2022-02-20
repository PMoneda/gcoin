package pow

import (
	"crypto/sha256"
	"fmt"

	"github.com/PMoneda/gcoin/block"
)

func Work(difficult int64, previousBlockHash []byte, blockData []byte) *block.Block {
	block := block.NewBlock(blockData, previousBlockHash)
	var nounce int64 = 0
	for {
		block.SetNonce(nounce)
		sum := sha256.Sum256(block.GetDataToSign())
		sum = sha256.Sum256(sum[:])
		if checkProofOfWork(difficult, sum) {
			block.SetHash(sum[:])
			break
		} else {
			nounce = nounce + 1
			if nounce%10000000 == 0 {
				fmt.Printf("keep trying %d\n", nounce)
			}
		}
	}
	return block
}

func checkProofOfWork(difficultDesired int64, sum [32]byte) bool {
	diffcultFound := int64(0)
	for _, ch := range sum {
		c := ch >> 4
		if c == 0 {
			diffcultFound = diffcultFound + 1
			if ch<<4 == 0 {
				diffcultFound = diffcultFound + 1
			} else {
				break
			}
		} else {
			break
		}
	}
	return diffcultFound == difficultDesired
}
