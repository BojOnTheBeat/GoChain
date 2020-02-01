package main

// Blockchain holds an ordered list of blocks
type Blockchain struct {
	blocks []*Block
}

// AddBlock appends a new block to the blockchain. This won't work initially because a new block chain is empty
// therefore we need a "Genesis Block" to start every new blockchain
func (bc *Blockchain) AddBlock(data string) {

	previousBlock := bc.blocks[len(bc.blocks)-1]

	newBlock := NewBlock(data, previousBlock.Hash)

	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock returns a new genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//NewBlockChain returns a new blockchain
func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
