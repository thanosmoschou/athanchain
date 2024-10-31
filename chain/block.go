/*
Author: Thanos Moschou
Description: This project is a simple implementation of
a Blockchain in Go. The goal is to create a basic
structure that simulates how a blockchain works,
allowing you to understand its fundamental components
and functionality.
*/

// Package chain contains some useful tools that help to implement the blockchain.
package chain

type Block struct {
	Index     int
	Timestamp string
	Data      []string
	PrevHash  string
	Hash      string
	Nonce     int
}

// NewBlock accepts the index of the previous block, the data of the previous block and its hash and
// returns a pointer to the new block. Now PoW added to the process.
func NewBlock(prevIndex int, data []string, prevHash string) *Block {
	currBlockPointer := &Block{}

	currBlockPointer.mine(prevIndex, data, prevHash)

	return currBlockPointer
}
