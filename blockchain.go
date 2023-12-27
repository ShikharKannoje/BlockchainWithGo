package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Block structure
type Block struct {
	hash         string                 // block hash
	previousHash string                 // hash of previous block
	timestamp    time.Time              // timestamp when the block was created
	pow          int                    // Proof of work, amount of effort taken to derive the current block's hash
	data         map[string]interface{} //data of the block
}

// Blockchain structure
type Blockchain struct {
	genesisBlock Block   //genesisBlock is the first block of blockchain
	chain        []Block // array of blocks
	difficulty   int     //difficulty is the minimum efforts required by miner to mine and include a block in blockchain
}

// Calculating the hash of a block
func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockdata := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockchain := sha256.Sum256([]byte(blockdata))
	return fmt.Sprintf("%x", blockchain)
}

// Mining new blocks
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

// creating the genesis block
func CreateBlockchain(difficutly int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}

	return Blockchain{
		genesisBlock: genesisBlock,
		chain:        []Block{genesisBlock},
		difficulty:   difficutly,
	}
}

// adding a new block
func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{ // creating a new data
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1] // figuring out the last block
	newBlock := Block{                   // creating a new block with data
		previousHash: lastBlock.hash,
		data:         blockData,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)         // mining of newblock to calculate hash and generate pow
	b.chain = append(b.chain, newBlock) // appending into the block chain
}

// checking the validity of the block
func (b *Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

// using the blockchain to make transaction

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	// record transactions on the blockchain for Tom, Brad, and John
	blockchain.addBlock("Tom", "John", 5)
	blockchain.addBlock("Brad", "John", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.isValid())
}
