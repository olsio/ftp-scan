package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type store struct {
	db *bolt.DB
}

func OpenDb(file string) *store {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := new(store)
	store.db = db
	return store
}
