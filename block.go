package main

import (
	"time"
)

// Block is The "Block" part of a Blockchain
type Block struct {
	Timestamp     int64
	Data          []byte // Actual transaction data
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// func (b *Block) setHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)

// 	b.Hash = hash[:]
// }

// NewBlock creates a block given data and the hash of a previous block
func NewBlock(data string, PrevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
