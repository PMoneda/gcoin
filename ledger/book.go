package ledger

import (
	"bytes"
	"errors"

	"github.com/PMoneda/gcoin/block"
	"github.com/PMoneda/gcoin/utils"
	bolt "go.etcd.io/bbolt"
)

type LedgerBook struct {
	db           *bolt.DB
	version      string
	powDifficult int64
	head         []byte
	origin       []byte
}

func (ledger *LedgerBook) GetHeadBlock() *block.Block {
	return ledger.GetBlock(ledger.head)
}

func (ledger *LedgerBook) GetOriginBlock() *block.Block {
	return ledger.GetBlock(ledger.origin)
}

func (ledger *LedgerBook) GetBlock(hash []byte) *block.Block {
	var savedBlock *block.Block
	ledger.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Blocks"))
		data := bucket.Get(hash)
		if data != nil {
			savedBlock = block.FromByteArray(data)
		}
		return nil
	})
	return savedBlock
}

func (ledger *LedgerBook) attachBlockToLedgerBook(_block *block.Block) error {
	return ledger.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Blocks"))
		meta := tx.Bucket([]byte("Metadata"))

		isFirstBlock := false
		origin := meta.Get([]byte("ORIGIN"))
		if bytes.Compare(origin, utils.Int32ToByteArrayNBytes(0, 32)) == 0 {
			meta.Put([]byte("ORIGIN"), _block.GetHash())
			ledger.origin = _block.GetHash()
			isFirstBlock = true
		}
		if !isFirstBlock && bytes.Compare(meta.Get([]byte("HEAD")), _block.GetPreviousHash()) != 0 {
			return errors.New("invalid block position: previous hash not match with HEAD")
		}
		if !block.CheckBlockConsistency(_block) {
			return errors.New("block consistency hashing cannot be validated")
		}
		if err := bucket.Put(_block.GetHash(), _block.ToBytes()); err != nil {
			return err
		}
		ledger.head = _block.GetHash()
		return meta.Put([]byte("HEAD"), _block.GetHash())
	})
}

func (ledger *LedgerBook) Close() error {
	return ledger.db.Close()
}
