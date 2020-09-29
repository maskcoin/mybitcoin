package main

import (
	"fmt"
	"time"
)

//1.定义结构（区块头的字段比正常的少）
//（1）前区块哈希
//（2）当前区块哈希
//（3）数据
//2.创建区块
//3.生成哈希
//4.引入区块链
//5.添加区块
//6.重构代码

func main() {
	bc := NewBlockChain()
	bc.AddBlock([]byte("班主任来了，大家快跑"))
	bc.AddBlock([]byte("班主任走了，大家鼓掌"))

	for i, block := range bc.Blocks {
		fmt.Printf("++++++++++++++++++++ %d ++++++++++++++++++++\n", i)
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
