package main

import (
	"chapter2/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("first block")
	chain.AddBlock("second block")
	chain.AddBlock("third block")

	for _, block := range chain.Blocks {
		fmt.Printf("Pre Hash: %x\n", block.PreHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)

		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
