package main

import (
    "fmt"
    "bytes"
    "crypto/sha256"
)

type Block struct {
    Hash []byte                 //Hash of this block
    PrevHash []byte
    Data []byte                 //Hash of prev block
}

type Chain struct{
    BlockLink []Block
}


func (block *Block) Hashify(){
    info := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
    hash := sha256.Sum256(info)
    blockHash := hash[:]
    block.Hash = blockHash
}


func NewBlock(data string, prevBlockHash []byte) Block{
    block := Block{Data: []byte(data), PrevHash: prevBlockHash}
    block.Hashify()

    return block
}


func (blockchain *Chain) NextBlock(data string) {
    prevBlock := blockchain.BlockLink [len(blockchain.BlockLink)-1]
    prevBlockHash := prevBlock.Hash
    new := NewBlock(data,[]byte(prevBlockHash))
    blockchain.BlockLink = append(blockchain.BlockLink,new)

}

func NewBlockchain() Chain {
    new := NewBlock("Genesis", []byte{})
    return Chain{[]Block{new}}
}


func main(){
    blockchain := NewBlockchain()

    blockchain.NextBlock("Hans to Kristoff: 60")
    blockchain.NextBlock("Kristoff to Anna: 40")
    blockchain.NextBlock("Anna to Elsa: 10")
    blockchain.NextBlock("Hans to Kristoff: 60")

    for _, showBlock := range blockchain.BlockLink {
        fmt.Printf("Previous Hash: %x\n", showBlock.PrevHash)
        fmt.Printf("Data in Block: %s\n", showBlock.Data)
        fmt.Printf("Hash: %x\n", showBlock.Hash)
    }

}