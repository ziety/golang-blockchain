package blockchain

// BlockChain represents a simple blockchain with a slice of blocks.
type BlockChain struct {
	Blocks []*Block
}

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

// AddBlock adds a new block to the blockchain.
func (chain *BlockChain) AddBlock(data string) {
	// Get the previous block in the blockchain.
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	// Create a new block with the given data and the previous block's hash.
	new := CreateBlock(data, prevBlock.Hash)

	// Append the new block to the blockchain.
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis creates the first block in the blockchain.
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initializes a new blockchain with the genesis block.
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
