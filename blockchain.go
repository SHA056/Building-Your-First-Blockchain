package main

import (
    "fmt"
    "bytes"
    "crypto/sha256"
)

type Block struct {
    // Hash string                 //Hash of this block
    PrevHash string
    Data string                 //Hash of prev block
}

type Chain struct{
    BlockLink []Block
}


func Hashify(pBlock Block) Block{
    info := bytes.Join([][]byte{pBlock.Data, pBlock.PrevHash}, []byte{})
    hash := sha256.Sum256(info)
    blockHash := hash[:]
    return blockHash
}


func NewBlock(data string, prevBlockHash Block) Block{
    block := Block{PrevHash: prevBlockHash, Data: data}
    return block
}



func (blockchain Chain) NextBlock(data string) {
    // if (len(blockchain.BlockLink)-1) == 0 {
    //     return CreateBlock("Genesis", []byte{})
    // }
    prevBlock := blockchain.BlockLink [len(blockchain.BlockLink)-1]
    prevBlockHash := Hashify(prevBlock)                 // Go through Hash function
    new := NewBlock(data,prevBlockHash)
    blockchain.BlockLink = append(blockchain.BlockLink,new)

}


func main(){
    blockchain := NextBlock("Genesis")

    blockchain.NextBlock("Hans to Kristoff: 60")
    blockchain.NextBlock("Kristoff to Anna: 40")
    blockchain.NextBlock("Anna to Elsa: 10")

    for _, showBlock := range blockchain.BlockLink {
        fmt.Printf("Previous Hash: ", showBlock.PrevHash)
        fmt.Printf("Data in Block: ", showBlock.Data)
        // fmt.Printf("Hash: ", showBlock.Hash)
    }

}