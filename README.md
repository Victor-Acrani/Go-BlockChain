# Go-BlockChain
Basic Golang framework for a blockchain. This repository is an exercise proposed by the Youtube channel [Tensor Programming](https://www.youtube.com/@TensorProgramming).

## Part 2
In part 2, is implemented the Proof of Work (PoW) algorithm. PoW is important to prevent fraud and network manipulation by requiring miners to solve complex mathematical problems before they can add new blocks to the blockchain.

To experiment with how the cost of PoW changes with problem difficulty, you can change the `difficulty` constant in the proof.go file. This constant means the amount of 0s needed to have a valid hash.

[Here you have an explanation of how PoW works.](https://creativedata.stream/bitcoin-proof-of-work/)

Link of Part 2 of this tutorial. [Building a Blockchain in Golang - Part 1](https://www.youtube.com/watch?v=aE4eDTUAE70&list=PLJbE2Yu2zumC5QE39TQHBLYJDB2gfFE5Q&index=2)

## How to clone
```sh
git clone "git@github.com:Victor-Acrani/Go-BlockChain.git"
```

## How to run 
```sh
go run main.go
```