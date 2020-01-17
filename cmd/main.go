package main

import "fmt"

func main() {
	chain := InitBlockchain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
