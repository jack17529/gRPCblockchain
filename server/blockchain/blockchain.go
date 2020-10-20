package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// this will be package blockchain

// because this is a blockchain package that is why we are doing it again.
// although it is already present in the proto file.
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// we only have block struct in proto file not this.
type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// we won't declare SHA encryption in this function but in its own function.
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}
	block.setHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

// it's like a CTOR but in go
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

// we need two branches as this is getting messed up now.
