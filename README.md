# Go-BlockChain
Basic Golang framework for a blockchain. This repository is an exercise proposed by the Youtube channel [Tensor Programming](https://www.youtube.com/@TensorProgramming).

## Part 1
In part 1, the `block` and `blockchain` structures are created. In this simples script it is possible to create a blockchain and add blocks into it.

 When a blockchain is created it is necessary to initialize it with a `genesis block` which is the initial block of the chain and it does not carry any relevant data as it is just a starting point.

To add a new block in the chain, is necessary to provide some data in order to create it and calculate its hash. Then the block is added by the blockchain.

Link of Part 1 of this tutorial. [Building a Blockchain in Golang - Part 1](https://www.youtube.com/watch?v=mYlHT9bB6OE&list=PLJbE2Yu2zumC5QE39TQHBLYJDB2gfFE5Q)

## How to clone
```sh
git clone "git@github.com:Victor-Acrani/Go-BlockChain.git"
```

## How run 
```sh
go run main.go
```