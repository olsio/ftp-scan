package store

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/olsio/ftp-scan/types"
)

const mainBucket = "target"

type Store struct {
	db *bolt.DB
}

func Open() *Store {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	store := new(Store)
	store.db = db
	return store
}

func (store *Store) Close() error {
	return store.db.Close()
}

func (store *Store) Clear() error {
	return store.db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(mainBucket))
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	})
}

func (store *Store) AddResult(target string, result types.Result) error {
	return store.db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(mainBucket))

		result, err := json.Marshal(result)

		if err != nil {
			log.Fatal(err)
			return err
		}

		return b.Put([]byte(target), result)
	})
}

func (store *Store) Contains(target string) (bool, error) {
	var result bool
	err := store.db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(mainBucket))
		result = bucket.Get([]byte(target)) != nil
		return err
	})
	return result, err
}

func (store *Store) List() error {
	return store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte([]byte(mainBucket)))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
}
