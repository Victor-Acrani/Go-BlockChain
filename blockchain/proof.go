package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

/*
--- Proof of work ---
- Take the data from the block
- Create a counter (nonce) which starts at 0
- Create a hash of the data plus the counter
- Check the hash of the data plus the counter
- Check the hash to see if it meets a set of requirements
- Requirements:
	The First few bytes must contains 0s
*/

// Difficulty gives a constant level of difficulty for minning
const Difficulty = 12

// Proof of Work structure
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// Create new proof of work.
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)                  // cast 1 to big int
	target.Lsh(target, uint(256-Difficulty)) // left shift bytes

	pow := &ProofOfWork{Block: b, Target: target}
	return pow
}

// Join block fields with difficulty and return the data in byte format.
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToBytes(int64(nonce)),
			ToBytes(int64(Difficulty))},
		[]byte{},
	)
	return data
}

// Parse int64 into bigEndian and return it in byte format.
func ToBytes(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// Calculate the nonce and the hash.
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		// the hash must be lower then the target to fulfil the requirement of the 0s in the beginning
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate block hash.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
