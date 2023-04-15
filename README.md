# Go-BlockChain
Basic Golang framework for a blockchain. This repository is an exercise proposed by the Youtube channel [Tensor Programming](https://www.youtube.com/@TensorProgramming).

## Part 3
In part 3, a persistence layer is added using BadgerDB and a CLI interface so that the user can interact with the application. Now each block of the chain is serialized and saved in a database and to recover the whole blockchain it is necessary to have only the last hash value. 

[BadgerDB](https://dgraph.io/docs/badger/get-started/)

Link of Part 3 of this tutorial. [Building a Blockchain in Golang - Part 3](https://www.youtube.com/watch?v=szOZ3p-5YIc&list=PLJbE2Yu2zumC5QE39TQHBLYJDB2gfFE5Q&index=3)

## How to clone
```sh
git clone "git@github.com:Victor-Acrani/Go-BlockChain.git"
```

## How to run 
Read usage:
```sh
go run main.go
```
Add block:
```sh
go run main.go add -block "send 15 BTC to John Doe"
```
Print blockchain:
```sh
go run main.go print
```