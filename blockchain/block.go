package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block represents a single block in the blockchain.
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock creates a new block with the given data and previous block's hash.
func CreateBlock(data string, prevHash []byte) *Block {
	// Create a new block with initial values.
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	// Generate a proof of work for the block and set its hash and nonce.
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis creates the first block in the blockchain.
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
