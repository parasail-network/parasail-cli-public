package main

import (
	"github.com/dgraph-io/badger/v4"
)

type LocalDB struct {
	db *badger.DB
}

func NewLocalDB(name string) (*LocalDB, error) {
	opts := badger.DefaultOptions("./badger")
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &LocalDB{db: db}, nil
}

func (l *LocalDB) Set(key string, value string) error {
	err := l.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
	return err
}

func (l *LocalDB) Get(key string) (*string, error) {
	var v string
	vPtr := &v
	err := l.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			*vPtr = string(val)
			return nil
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return vPtr, nil
}

func (l *LocalDB) Delete(key string) error {
	return l.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}
