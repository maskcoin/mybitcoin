package main

import (
	"time"
)

type Block struct {
	Version       uint64 //区块版本号
	PrevBlockHash []byte //前区块哈希
	MerkleRoot    []byte //先填写为空，后续v4的时候使用
	TimeStamp     uint64 //从1970.1.1至今的秒数
	Difficulty    uint64 //挖矿的难度值，v2时使用
	Nonce         uint64 //随机数，挖矿找的就是它
	Data          []byte
	Hash          []byte //当前区块hash，区块中本不存在，为了方便我们添加进来
}

//创建区块，对Block的每一个字段填充数据即可
func NewBlock(data []byte, prevBlockHash []byte) *Block {
	block := &Block{
		Version:       0,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    nil,
		TimeStamp:     uint64(time.Now().Unix()),
		Difficulty:    Bits,
		Data:          data,
		Hash:          nil,
	}

	//block.SetHash()
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()

	block.Nonce = nonce
	block.Hash = hash

	return block
}

