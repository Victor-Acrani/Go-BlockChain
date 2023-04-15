package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Calculate block hash.
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Create new block.
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash}
	block.DeriveHash()
	return block
}

// Create genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
