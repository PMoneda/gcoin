package utils

import (
	"fmt"

	"go.etcd.io/bbolt"
)

func OpenBoltDb(name string) (*bbolt.DB, error) {
	return bbolt.Open(fmt.Sprintf("./%s.db", name), 0666, nil)
}

func SaveBoltData(db *bbolt.DB, bucketName string, key string, value string) error {
	tx, err := db.Begin(true)
	if err != nil {
		tx.Rollback()
		return err
	}
	bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
	if err != nil {
		tx.Rollback()
		return err
	}
	err = bucket.Put([]byte(key), []byte(value))
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func GetBoltData(db *bbolt.DB, bucketName string, key string) (string, error) {
	var value []byte
	db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		value = bucket.Get([]byte(key))
		return nil
	})
	return string(value), nil
}

func ListBoltBucket(db *bbolt.DB, bucketName string) ([]string, error) {
	contactList := make([]string, 0)
	db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		bucket.ForEach(func(k, v []byte) error {
			contactList = append(contactList, fmt.Sprintf("%s:%s", string(k), string(v)))
			return nil
		})
		return nil
	})
	return contactList, nil
}

func ListBoltBucketKeys(db *bbolt.DB, bucketName string) ([]string, error) {
	list := make([]string, 0)
	db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return nil
		}
		bucket.ForEach(func(k, v []byte) error {
			list = append(list, string(k))
			return nil
		})
		return nil
	})
	return list, nil
}

func DeleteBoltData(db *bbolt.DB, bucketName string, key string) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	bucket := tx.Bucket([]byte(bucketName))
	err = bucket.Delete([]byte(key))
	if err != nil {
		return err
	}
	return tx.Commit()
}
