package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// define block
type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PreHash      []byte
	Nonce        int
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

// create new block
func CreateBlock(txs []*Transaction, prehash []byte) *Block {
	newBlock := &Block{[]byte{}, txs, prehash, 0}
	pow := NewProof(newBlock)
	nonce, hash := pow.Run()

	newBlock.Nonce = nonce
	newBlock.Hash = hash

	return newBlock
}

// Init block
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
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
