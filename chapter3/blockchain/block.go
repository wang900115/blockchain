package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// define block
type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
	Nonce   int
}

// create new block
func CreateBlock(data string, prehash []byte) *Block {
	newBlock := &Block{[]byte{}, []byte(data), prehash, 0}
	pow := NewProof(newBlock)
	nonce, hash := pow.Run()

	newBlock.Nonce = nonce
	newBlock.Hash = hash

	return newBlock
}

// Init block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// serialize
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)
	return res.Bytes()
}

// deserialize
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

// handle error
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
