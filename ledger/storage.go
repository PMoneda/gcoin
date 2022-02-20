package ledger

import (
	"errors"
	"os"

	"github.com/PMoneda/bblock/config"
	"github.com/PMoneda/bblock/utils"
	bolt "go.etcd.io/bbolt"
)

func NewLedgerFromPath(path string) (*LedgerBook, error) {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		if b, err := tx.CreateBucket([]byte("Metadata")); err != nil {
			return err
		} else if err := b.Put([]byte("HEAD"), utils.Int32ToByteArrayNBytes(0, 32)); err != nil {
			return err
		} else if err := b.Put([]byte("ORIGIN"), utils.Int32ToByteArrayNBytes(0, 32)); err != nil {
			return err
		} else if err := b.Put([]byte("Version"), []byte(config.Version)); err != nil {
			return err
		} else if err := b.Put([]byte("PoW_Difficult"), utils.Int64ToByteArray(config.PoW_Difficult)); err != nil {
			return err
		}
		if _, err := tx.CreateBucket([]byte("Blocks")); err != nil {
			return err
		}
		return nil
	})
	return &LedgerBook{
		db:           db,
		version:      config.Version,
		head:         utils.Int32ToByteArrayNBytes(0, 32),
		origin:       utils.Int32ToByteArrayNBytes(0, 32),
		powDifficult: config.PoW_Difficult,
	}, nil
}

func OpenLedger(path string) (*LedgerBook, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return NewLedgerFromPath(path)
	}
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return nil, err
	}

	ledger := &LedgerBook{
		db: db,
	}
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Metadata"))
		ledger.head = bucket.Get([]byte("HEAD"))
		ledger.origin = bucket.Get([]byte("ORIGIN"))
		ledger.version = string(bucket.Get([]byte("Version")))
		ledger.powDifficult = utils.Int64FromArray(bucket.Get([]byte("PoW_Difficult")))
		return nil
	})
	return ledger, nil
}
