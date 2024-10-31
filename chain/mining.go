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

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
)

// This constant tells how many 0's should be at the beginning of each hash.
const Difficulty int = 4

func (block *Block) mine(prevIndex int, data []string, prevHash string) {
	newIndex := prevIndex + 1
	newTimestamp := time.Now().Format(time.RFC3339)
	dataAsString := strings.Join(data, "-")
	nonce := 0
	prefix := strings.Repeat("0", Difficulty)

	for {
		// Hash the whole block's data, including index, timestamp and prevHash
		record := string(newIndex) + newTimestamp + dataAsString + prevHash + string(nonce)
		h := sha256.New()
		h.Write([]byte(record))
		currHash := hex.EncodeToString(h.Sum(nil))

		if strings.HasPrefix(currHash, prefix) {
			block.Index = newIndex
			block.Timestamp = newTimestamp
			block.Data = data
			block.PrevHash = prevHash
			block.Hash = currHash
			block.Nonce = nonce
			return
		} else {
			nonce++
		}
	}

}
