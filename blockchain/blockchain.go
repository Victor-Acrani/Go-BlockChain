package blockchain

// Blockchain structure
type BlockChain struct {
	Blocks []*Block
}

// Append block to blockchain.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Create a new blockchain.
func InitBlockChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{Genesis()}}
}
