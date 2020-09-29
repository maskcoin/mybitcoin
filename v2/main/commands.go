package main

import (
	"fmt"
	"time"
)

//实现具体的命令
func (cli *CLI) AddBlock(data []byte) {
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块成功!\n")
}

func (cli *CLI) PrintChain()  {
	it := cli.bc.NewIterator()

	for it.current != nil {
		block := it.Next()
		fmt.Printf("++++++++++++++++++++++++++++++++++++++++\n")
		fmt.Printf("Version: %d\n", block.Version)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("MerkleRoot: %x\n", block.MerkleRoot)
		timeFormatStr := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("TimeStamp: %s\n", timeFormatStr)
		fmt.Printf("Difficulty: %d\n", block.Difficulty)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProofOfWork(block)

		fmt.Printf("IsValid: %v\n", pow.IsValid())
	}
}
