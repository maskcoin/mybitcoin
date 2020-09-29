package main

import (
	"bytes"
	"encoding/gob"
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

//序列化，将区块转换为字节切片
func (block *Block) Serialize() (ret []byte) {
	var buffer bytes.Buffer
	//定义编码器
	//间接赋值是指针存在的最大意义
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(block)
	if err != nil {
		panic(err)
	}
	ret = buffer.Bytes()

	return
}

func Deserialize(data []byte) *Block {
	//fmt.Printf("解码传入的数据：%x\n", data)

	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	//间接赋值是指针存在的最大意义
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}

