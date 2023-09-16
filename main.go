package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain represents a simple blockchain with a slice of blocks.
type BlockChain struct {
	blocks []*Block
}

// Block represents a single block in the blockchain.
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash calculates the hash of the block based on its data and the previous block's hash.
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a new block with the given data and previous block's hash.
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock adds a new block to the blockchain.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Genesis creates the first block in the blockchain.
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initializes a new blockchain with the genesis block.
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	// Adding some sample blocks to the blockchain.
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// Printing information about each block in the blockchain.
	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
