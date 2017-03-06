package store

import (
	"encoding/binary"
	"log"

	"github.com/boltdb/bolt"
)

type Store struct {
	db *bolt.DB
}

func NewStore(file string) *Store {
	db, err := bolt.Open(file, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := new(Store)
	store.db = db
	return store
}

func (store *Store) AddScanTarget(target []byte) error {
	return store.db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b, _ := tx.CreateBucketIfNotExists([]byte("target"))
		id, _ := b.NextSequence()

		// Persist bytes to users bucket.
		return b.Put(itob(id), target)
	})
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
