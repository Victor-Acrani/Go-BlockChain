package blockchain

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

// Create genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
