package block

import (
	"bytes"
	"crypto/sha256"

	"github.com/PMoneda/gcoin/utils"
)

type Block struct {
	nonce        int64
	body         []byte
	previousHash []byte
	hash         []byte
}

const fields = 3
const int64Size = 8
const sep = "|||"

func (block *Block) GetDataToSign() []byte {
	data := make([]byte, len(block.body)+len(block.previousHash)+int64Size+(fields*len(sep)))
	offset := 0
	i := 0
	i = copy(data[offset:], block.body)
	offset = offset + i
	i = copy(data[offset:], []byte(sep))
	offset = offset + i
	i = copy(data[offset:], block.previousHash)
	offset = offset + i
	i = copy(data[offset:], []byte(sep))
	offset = offset + i
	b := utils.Int64ToByteArray(block.nonce)
	i = copy(data[offset:], b)
	offset = offset + i
	copy(data[offset:], []byte(sep))
	return data
}

func (block *Block) SetNonce(nounce int64) {
	block.nonce = nounce
}

func (block *Block) GetNounceOffset() int {
	return len(block.body) + len(block.previousHash) + ((fields - 1) * len(sep))
}

func (block *Block) SetHash(hash []byte) {
	block.hash = hash
}

func (block *Block) ToBytes() []byte {
	data := block.GetDataToSign()
	finalData := make([]byte, len(data)+len(block.hash))
	i := copy(finalData[0:], data)
	copy(finalData[i:], block.hash)
	return finalData
}

func (block *Block) GetHash() []byte {
	return block.hash
}

func (block *Block) GetPreviousHash() []byte {
	return block.previousHash
}

func (block *Block) GetBody() []byte {
	return block.body
}

func NewBlock(body []byte, previous []byte) *Block {
	return &Block{
		nonce:        0,
		body:         body,
		previousHash: previous,
	}
}

func FromByteArray(array []byte) *Block {
	block := &Block{}
	parts := bytes.Split(array, []byte(sep))
	block.body = parts[0]
	block.previousHash = parts[1]
	block.nonce = utils.Int64FromArray(parts[2])
	block.hash = parts[3]
	return block
}

func CheckBlockConsistency(block *Block) bool {
	sum := sha256.Sum256(block.GetDataToSign())
	return bytes.Compare(sum[:], block.hash) == 0
}
