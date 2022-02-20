package pow

import (
	"testing"

	"github.com/PMoneda/bblock/config"
	"github.com/PMoneda/bblock/utils"
)

func Test_ShouldCalculateHash(t *testing.T) {
	previousHash := utils.Int32ToByteArrayNBytes(0, 32)
	data := []byte("Hello Block")
	block := Work(config.PoW_Difficult, previousHash, data)
	hash := block.GetHash()
	if hash[0] != 0 || hash[1] != 0 || hash[2] >= 8 {
		t.Fail()
	}

}
