package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

/*
--- LINKS ---
https://pkg.go.dev/github.com/dgraph-io/badger/v4#Item.Value
https://dgraph.io/docs/badger/get-started/
*/

const (
	dbPath = "./tmp/blocks"
	lasthashIndex = "lh"
)

// Blockchain structure
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

// BlockChainIterator structure
type BlockStructureIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// Create a new blockchain.
func InitBlockChain() *BlockChain {
	var lastHash []byte
	// set database default options with provided path
	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	HandleError(err)

	// create ReadWrite transaction with BadgerDB
	// Retrieve blockchain if exists otherwise creates one
	db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(lasthashIndex))
		if err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()

			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			HandleError(err)

			err = txn.Set([]byte(lasthashIndex), genesis.Hash)
			lastHash = genesis.Hash
			return err
		} else {
			item, err := txn.Get([]byte(lasthashIndex))
			HandleError(err)
			lastHash, err = item.ValueCopy(nil)
			return err
		}
	})

	HandleError(err)

	blockchain := BlockChain{LastHash: lastHash, Database: db}
	return &blockchain
}

// Append block to blockchain.
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	// create ReadOnly transaction with BadgerDB
	// get last hash value
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(lasthashIndex))
		HandleError(err)
		lastHash, err = item.ValueCopy(nil)
		return err
	})

	HandleError(err)
	newBlock := CreateBlock(data, lastHash)

	// create ReadWrite transaction with BadgerDB
	// Save new block and update last hash value
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		HandleError(err)
		err = txn.Set([]byte(lasthashIndex), newBlock.Hash)

		chain.LastHash = newBlock.Hash
		return err
	})
	HandleError(err)
}

// Gives an Iterator for a Blockchain
func (chain *BlockChain) Iterator() *BlockStructureIterator {
	var iterator = &BlockStructureIterator{
		CurrentHash: chain.LastHash,
		Database:    chain.Database,
	}

	return iterator
}

// Gives next element of an iterator
func (iter *BlockStructureIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		HandleError(err)
		encodedBlock, err := item.ValueCopy(nil)
		block = Deserialize(encodedBlock)

		return err
	})
	HandleError(err)

	iter.CurrentHash = block.PrevHash
	return block
}
