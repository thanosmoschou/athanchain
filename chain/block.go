/*
Author: Thanos Moschou
Description: This project is a simple implementation of
a Blockchain in Go. The goal is to create a basic
structure that simulates how a blockchain works,
allowing you to understand its fundamental components
and functionality.
*/

package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      []string
	PrevHash  string
	Hash      string
}

func NewBlock(prevIndex int, data []string, prevHash string) *Block {
	newIndex := prevIndex + 1
	newTimestamp := time.Now().Format(time.RFC3339)
	dataAsString := strings.Join(data, "-")

	// Hash the whole block's data, including index, timestamp and prevHash
	record := string(newIndex) + newTimestamp + dataAsString + prevHash

	h := sha256.New()
	h.Write([]byte(record))
	currHash := hex.EncodeToString(h.Sum(nil))

	return &Block{
		Index:     newIndex,
		Timestamp: newTimestamp,
		Data:      data,
		PrevHash:  prevHash,
		Hash:      currHash,
	}
}
