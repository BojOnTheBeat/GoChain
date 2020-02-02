package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// In Bitcoin targetBits is the block header that stores the
// difficulty at which the block was mined. In the real world
// this is couple with a target adjusting algorithm but we won't
// do that now so we define the difficulty as a global constant
const targetBits = 24

// ProofOfWork is a struct that holds a block and the target hash
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork Calculates the target and returns a ProofOfWork struct
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)

	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow

}

// PerpareData returns the data we want to hash
func (pow *ProofOfWork) PerpareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data

}

// IntToHex converts an integer to a hex
// TODO: maybe change this to bigendian bytes array?
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// func IntToHex(n int64) []byte {
// 	return []byte(strconv.FormatInt(n, 16))
// }

// Run in a loop:
// 1. PrepareData
// 2. Hash it with Sha-256
// 3. convert the hash to a big integer
// 4. Compare the integer with the target.
// (We do this till we get a hash that satisfies the target.)
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.PerpareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}
