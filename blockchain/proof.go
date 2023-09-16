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

// Difficulty determines the difficulty level for mining new blocks.
const Difficulty = 18

// ProofOfWork represents the proof of work mechanism for mining blocks.
type ProofOfWork struct {
	Block  *Block   // The block to be mined.
	Target *big.Int // The target value that the hash must be less than.
}

// NewProof creates a new proof of work instance for a given block.
func NewProof(b *Block) *ProofOfWork {
	// Create a target hash that starts with a certain number of leading zeros based on Difficulty.
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// InitData prepares the data for mining by including nonce and difficulty.
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

// Run performs the mining process and returns the nonce and resulting hash.
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	// Continue mining until a hash is found that meets the target criteria.
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate checks if a mined block's hash meets the target criteria.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// ToHex converts an integer to a byte slice.
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
