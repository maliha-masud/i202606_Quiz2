package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Hash      string
	PrevHash  string
	Timestamp string
	Data      string
}

func calculateHash(index int, timestamp, data, prevHash string) string { //calculate block hash using index, timestamp, data, and previous hash
	hashInput := fmt.Sprintf("%d%s%s%s", index, timestamp, data, prevHash)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}

func DisplayAllBlocks(chain []Block) { //display all blocks in the blockchain
	fmt.Println("Display All Blocks:")
	for _, block := range chain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
	}
}

func NewBlock(prevBlock Block, data string) Block { //create new block in the blockchain
	newIndex := prevBlock.Index + 1
	newTimestamp := time.Now().String()
	newHash := calculateHash(newIndex, newTimestamp, data, prevBlock.Hash)

	return Block{
		Index:     newIndex,
		Hash:      newHash,
		PrevHash:  prevBlock.Hash,
		Timestamp: newTimestamp,
		Data:      data,
	}
}

func ModifyBlock(block *Block, newData string) { //update block's data
	block.Data = newData
	block.Timestamp = time.Now().String()
	block.Hash = calculateHash(block.Index, block.Timestamp, block.Data, block.PrevHash)
}

func main() {
	genesisBlock := Block{ //create genesis block (initial block)
		Index:     0,
		PrevHash:  "",
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
	}

	blockchain := []Block{genesisBlock}

	block1 := NewBlock(blockchain[len(blockchain)-1], "Data for Block 1")
	blockchain = append(blockchain, block1)

	block2 := NewBlock(blockchain[len(blockchain)-1], "Data for Block 2")
	blockchain = append(blockchain, block2)

	DisplayAllBlocks(blockchain)
	ModifyBlock(&blockchain[1], "Updated data for Block 2")
	DisplayAllBlocks(blockchain)
}
