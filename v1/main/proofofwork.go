package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const Bits  = 20

type ProofOfWork struct {
	block *Block
	target *big.Int //系统提供的，是固定的
}

func NewProofOfWork(block *Block) (pow *ProofOfWork) {
	pow = &ProofOfWork{
		block: block,
	}

	//固定的难度值
	//写难度值，难度值应该是推导出来的，我们为了简化，把难度值先写成固定的，一切完成之后再去推导
	//000012fa9b916eb9078f8d98a7864e697ae83ed54f5146bd84452cdafd043c19
	//targetStr := "000012fa9b916eb9078f8d98a7864e697ae83ed54f5146bd84452cdafd043c19"
	//
	//var tmpBigInt big.Int
	//pow.target, _ = tmpBigInt.SetString(targetStr, 16)

	//程序推导难度值，推导前导为3个0的难度值
	//000012fa9b916eb9078f8d98a7864e697ae83ed54f5146bd84452cdafd043c19
	//初始化
	bigIntTmp := big.NewInt(1)
	//bigIntTmp.Lsh(bigIntTmp, 256)
	//bigIntTmp.Rsh(bigIntTmp, 20)
	bigIntTmp.Lsh(bigIntTmp, 256 - Bits)

	pow.target = bigIntTmp

	return
}

//这是pow的运算函数，为了获取挖矿的nonce值，同时返回区块的哈希值
func (pow *ProofOfWork) Run() (hash []byte, nonce uint64) {
	//1.获取block数据
	//block := pow.block
	for  {
		//2.拼接nonce
		data := pow.prepareData(nonce)
		//3.sha256
		hashArr := sha256.Sum256(data)
		hash = hashArr[:]
		//4.与难度值比较
		var bigIntTmp big.Int
		bi := bigIntTmp.SetBytes(hash)

		if bi.Cmp(pow.target) < 0 {
			fmt.Printf("挖矿成功! nonce:%d, 哈希值:%x\n", nonce, hash)
			break
		} else {
			nonce++
		}
	}

	return
}

func (pow *ProofOfWork) prepareData(nonce uint64) (data []byte) {
	block := pow.block
	data = bytes.Join([][]byte{
		Uint64ToByteSlice(block.Version),
		block.MerkleRoot,
		Uint64ToByteSlice(block.TimeStamp),
		Uint64ToByteSlice(block.Difficulty),
		block.Data,
		Uint64ToByteSlice(nonce),
	}, nil)

	return
}

func (pow *ProofOfWork) IsValid() bool {
	data := pow.prepareData(pow.block.Nonce)
	hashArr := sha256.Sum256(data)
	hash := hashArr[:]

	var bigIntTmp big.Int
	bi := bigIntTmp.SetBytes(hash)

	return bi.Cmp(pow.target) < 0
}
