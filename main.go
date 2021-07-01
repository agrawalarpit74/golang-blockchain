package main

import (
	"fmt"
	"github.com/agrawalarpit74/golang-blockchain/blockchain"
)

func main()  {

	chain := InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _,block := range chain.blocks {
		fmt.Printf("Previous hash: %X\n", block.PrevHash)
		fmt.Printf("Data in block %s\n", block.Data)
		fmt.Printf("Hash: %X\n", block.Hash)
		fmt.Println("===============")
	}

}
