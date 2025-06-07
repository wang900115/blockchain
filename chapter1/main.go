package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// define block
type Block struct {
	Hash    []byte
	PreHash []byte
	Data    []byte
}

// define blockchain( array of block )
type BlockChain struct {
	Blocks []*Block
}

// signature hash to new block
func (b *Block) deriveHash() {
	data := bytes.Join([][]byte{b.Data, b.PreHash}, []byte{})
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}

// create new block
func CreateBlock(data string, prehash []byte) *Block {
	newBlock := &Block{[]byte{}, prehash, []byte(data)}
	newBlock.deriveHash()
	return newBlock
}

// add existing block
func (bc *BlockChain) AddBlock(data string) {
	lastHash := bc.Blocks[len(bc.Blocks)-1].Hash
	newBlock := CreateBlock(data, lastHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Init block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Init blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {

	chain := InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {

		fmt.Printf("PreHash: %x\n", block.PreHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}

}
