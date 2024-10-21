# Dummy Blockchain in Go

This project is a simple implementation of a Blockchain in Go. The goal is to create a basic structure that simulates how a blockchain works, allowing you to understand its fundamental components and functionality.

## Project Overview
In this project, we implement a minimal blockchain where:

Each block contains an **index**, **timestamp**, **data**, **hash**, and the **hash of the previous block**.
The first block is called the **Genesis** Block, and **each new block links to the previous one**, forming a chain.
The project can be extended with mechanisms like **Proof of Work** for difficulty in mining blocks.

## Key Components

### Block Structure:

**Index**: A unique number identifying the block in the chain. <br>
**Timestamp**: The time the block was created. <br>
**Data**: The content stored in the block (e.g., messages or transactions). <br>
**PrevHash**: The hash of the previous block, used to maintain the chain. <br>
**Hash**: The hash of the current block, calculated using its contents. <br>

### Blockchain Structure:

A series of blocks where each block points to the previous one via PrevHash.
The blockchain starts with a Genesis Block.

### Hashing:

Each block's hash is calculated using the SHA-256 algorithm based on its index, timestamp, data, and the previous block's hash.

### Optional Features (for extended learning):

**Proof of Work (PoW)**: A mechanism where miners need to solve a computational challenge to add a block.

**Blockchain Validation**: Functions that ensure the chain is valid by verifying the hashes and links between blocks.

## Getting Started

### Prerequisites
You need to have Go installed on your machine to run this project.

### Installation

1. Clone the repository:

```bash
git clone https://github.com/thanosmoschou/athanchain.git
cd athanchain
```

2. Run the Go application:

```go
go run main.go
```

### Example Output

Running the project will output the blockchain with two blocks added after the Genesis Block

```yaml
Index: 0
Timestamp: 2024-10-21 12:34:56
Data: Genesis Block
Prev Hash: 
Hash: 0a3c9fa0b3d8c2322dd7459d912f477848b...

Index: 1
Timestamp: 2024-10-21 12:35:10
Data: Block 1 Data
Prev Hash: 0a3c9fa0b3d8c2322dd7459d912f477848b...
Hash: 9b7e4fcebb59c1d2e2f3a521ea57e456789...

Index: 2
Timestamp: 2024-10-21 12:35:25
Data: Block 2 Data
Prev Hash: 9b7e4fcebb59c1d2e2f3a521ea57e456789...
Hash: 7a4e9df3ab123b5675c2f55f89bbf9d5dc9...
```

### How it Works

**Genesis Block**: The first block in the blockchain is created manually.

**Adding New Blocks**: For each new block, we take the hash of the previous block and use it to link the new one.

**Hash Calculation**: The SHA-256 algorithm is used to generate a unique hash for each block, based on its contents.