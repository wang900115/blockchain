package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// define block
type Block struct {
	Timestamp    int64
	Hash         []byte
	Transactions []*Transaction
	PreHash      []byte
	Nonce        int
	Height       int
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Serialize())
	}

	tree := NewMerkleTree(txHashes)
	return tree.RootNode.Data
}

// create new block
func CreateBlock(txs []*Transaction, prehash []byte, height int) *Block {
	newBlock := &Block{time.Now().Unix(), []byte{}, txs, prehash, 0, height}
	pow := NewProof(newBlock)
	nonce, hash := pow.Run()

	newBlock.Nonce = nonce
	newBlock.Hash = hash

	return newBlock
}

// Init block
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{}, 0)
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
