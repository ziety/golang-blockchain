package main

import (
	"fmt"
	"strconv"

	"github.com/ziety/golang-blockchain/blockchain"
)

func main() {
	// Initialize the blockchain.
	chain := blockchain.InitBlockChain()

	// Adding some sample blocks to the blockchain.
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// Printing information about each block in the blockchain.
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		// Create a new proof of work for the block.
		pow := blockchain.NewProof(block)

		// Validate the proof of work.
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
