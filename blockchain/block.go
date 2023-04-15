package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Create new block.
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Nonce: 0}

	// hash is definded by proof of work
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}

// Create genesis block.
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Gives a serialized block.
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	HandleError(err)
	return res.Bytes()
}

// Retrieve a block from serialized data.
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	HandleError(err)
	return &block
}

// Handle error.
func HandleError(err error){
	if err != nil{
		log.Panic(err)
	}
}
