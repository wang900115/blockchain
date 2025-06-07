package blockchain

// define block
type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
	Nonce   int
}

// define blockchain( array of block )
type BlockChain struct {
	Blocks []*Block
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
