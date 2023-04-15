package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/Victor-Acrani/Go-BlockChain/blockchain"
)

// CommandLine structure
type CommandLine struct {
	blockchain *blockchain.BlockChain
}

// Shows the usage of the cli commands.
func (cli *CommandLine) printUsage() {
	fmt.Println("##################################################")
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
	fmt.Println("##################################################")
}

// Validate input arguments.
func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

// Add block in the chain.
func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added block!")
}

// Print blockchain in cli terminal.
func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next()
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		// create proof of work
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

// Execute cli menu function.
func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.HandleError(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.HandleError(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	defer os.Exit(0)
	// create block chain
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := CommandLine{blockchain: chain}
	cli.run()
}
