package main

import (
	"bytes"
	"crypto/sha256"
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
	Transactions	[]*Transaction
	Hash          []byte //当前区块hash，区块中本不存在，为了方便我们添加进来
}

//创建区块，对Block的每一个字段填充数据即可
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Version:       0,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    nil,
		TimeStamp:     uint64(time.Now().Unix()),
		Difficulty:    Bits,
		Transactions:          txs,
		Hash:          nil,
	}

	block.HashTransactions()

	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()

	block.Nonce = nonce
	block.Hash = hash

	return block
}

//模拟梅克尔根，做简单的处理
func (block *Block) HashTransactions() {
	//我们的交易的id就是交易的哈希值，所以我们可以将交易id拼接起来，整体做一个hash，作为MerkleRoot
	var hashes []byte
	for _, v := range block.Transactions {
		txid := v.TXID
		hashes = append(hashes, txid...)
	}

	hash := sha256.Sum256(hashes)

	block.MerkleRoot = hash[:]
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
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	//间接赋值是指针存在的最大意义
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}

