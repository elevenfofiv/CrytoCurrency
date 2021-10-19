package main

import (
	"fmt"

	"github.com/elevenfofiv/crytocurrency/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PreHash: %s\n", block.PrevHash)

	}
}
