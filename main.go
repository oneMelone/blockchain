package main

import (
	"fmt"

	"simplified-blockchain/blockchainDs"
	"strconv"
)

func main() {
	bc := blockchainDs.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.GetWholeChain() {
		fmt.Printf("Prev. hash :%x \n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchainDs.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}