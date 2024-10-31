/*
Author: Thanos Moschou
Description: This project is a simple implementation of
a Blockchain in Go. The goal is to create a basic
structure that simulates how a blockchain works,
allowing you to understand its fundamental components
and functionality.
*/

package main

import (
	"athanchain/chain"
	"fmt"
)

const blockTransactionLimit int = 3

// Careful with package variables...
var transCtr int = 0

func main() {
	var blockchain []*chain.Block = make([]*chain.Block, 0)

	// Create Genesis Block
	genesisData := []string{"Genesis Block... No other transactions here..."}
	genesisBlock := chain.NewBlock(-1, genesisData, "0")
	prevHash := genesisBlock.Hash
	prevIndex := genesisBlock.Index
	blockchain = append(blockchain, genesisBlock)

	var blockData []string = make([]string, 0)
	var cont bool = true

	for cont {
		transCtr++
		blockData = append(blockData, makeNewTransaction())

		if transCtr%blockTransactionLimit == 0 {
			newBlock := chain.NewBlock(prevIndex, blockData, prevHash)
			blockchain = append(blockchain, newBlock)

			printBlockchain(blockchain)

			prevHash = newBlock.Hash
			prevIndex = newBlock.Index
			transCtr = 0
			blockData = make([]string, 0)
		}

		cont = askToContinue()
		if !cont {
			// add the remaining data to a new block before you finish.
			newBlock := chain.NewBlock(prevIndex, blockData, prevHash)
			blockchain = append(blockchain, newBlock)

			printBlockchain(blockchain)
		}
	}

}

func askToContinue() bool {
	var decision rune
	fmt.Println("Continue? (Y/N)")
	fmt.Scanf("%c\n", &decision)
	return decision == 'Y' || decision == 'y'
}

func printBlockchain(blockchain []*chain.Block) {

	fmt.Println("******CURRENT BLOCKCHAIN******")

	for _, block := range blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

func makeNewTransaction() string {
	var senderName string
	var receiverName string
	var amount int

	fmt.Printf("Transaction No%d.\n", transCtr)

	fmt.Println("Sender: ")
	fmt.Scanf("%s\n", &senderName)

	fmt.Println("Receiver: ")
	fmt.Scanf("%s\n", &receiverName)

	fmt.Println("Amount: ")
	fmt.Scanf("%d\n", &amount)

	return fmt.Sprintf("%s sent to %s %v athancoins.", senderName, receiverName, amount)
}
