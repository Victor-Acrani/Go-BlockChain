package main

import (
	"fmt"

	"github.com/Victor-Acrani/Go-BlockChain/blockchain"
)

func main() {
	// create block chain
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First block after genesis")
	chain.AddBlock("Second block after genesis")
	chain.AddBlock("Third block after genesis")

	// iterate over blockchain
	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
