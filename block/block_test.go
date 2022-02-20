package block

import (
	"strings"
	"testing"
)

func Test_ShouldReturnASliceWithInformationToBeSigned(t *testing.T) {
	block := NewBlock([]byte("Hello World"), []byte("1111"))
	str := string(block.GetDataToSign())
	if !strings.HasPrefix(str, "Hello World|||1111|||") {
		t.Fail()
	}
}

func Test_ShouldReturnFullBytes(t *testing.T) {
	block := NewBlock([]byte("Hello World"), []byte("1111"))
	block.SetHash([]byte("123"))
	str := string(block.ToBytes())
	if !(strings.HasPrefix(str, "Hello World|||1111|||") && strings.HasSuffix(str, "|||123")) {
		t.Fail()
	}
}

func Test_ShouldReturnNounceOffset(t *testing.T) {
	block := NewBlock([]byte("Hello World"), []byte("1111"))
	offset := block.GetNounceOffset()
	cmp := len("Hello World|||1111|||")
	if offset != cmp {
		t.Fail()
	}
}

func Test_ShouldConvertToArrayAndConvertBackToBlock(t *testing.T) {
	block := NewBlock([]byte("Hello World"), []byte("1111"))
	block.SetNonce(123)
	block.SetHash([]byte("123"))
	byteArray := block.ToBytes()
	cloneBlock := FromByteArray(byteArray)
	if string(block.body) != string(cloneBlock.body) {
		t.Fail()
	}
	if string(block.previousHash) != string(cloneBlock.previousHash) {
		t.Fail()
	}
	if string(block.hash) != string(cloneBlock.hash) {
		t.Fail()
	}
	if block.nonce != cloneBlock.nonce {
		t.Fail()
	}

}
