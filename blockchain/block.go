package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}
type Block struct {
	Hash	[]byte
	Data    []byte
	PrevHash []byte
}

func (b *Block) DeriveHash()  {
	info := bytes.Join([][]byte{b.Data,b.PrevHash},[]byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string,prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.blocks[len(chain.blocks)-1]
	block:= CreateBlock(data,prevBlock.Hash)
	chain.blocks = append(chain.blocks, block)
}

func Genesis() *Block {
	return CreateBlock("Genesis",[]byte{})
}

func InitBlockChain() *BlockChain  {
	return &BlockChain{[]*Block{Genesis()}}
}
