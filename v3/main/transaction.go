package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

//定义交易结构
//定义input
//定义output
//设置交易ID

type TXInput struct {
	TXID    []byte //交易ID
	Index   int64  //output的索引
	Address string //解锁脚本，先使用地址来模拟
}

type TXOutput struct {
	Value   float64 //转账金额
	Address string  //锁定脚本
}

type Transaction struct {
	TXID      []byte
	TXInputs  []TXInput
	TXOutPuts []TXOutput
}

func (tx *Transaction) SetTXID() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)

	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]
}

//实现挖矿交易
//特点：只有输出，没有有效输入
func NewCoinbaseTx(address string) (tx *Transaction) {
	inputs := []TXInput{
		{
			TXID:    nil,
			Index:   -1,
			Address: "genesisInfo",
		},
	}

	outputs := []TXOutput{
		{
			Value:   12.5,
			Address: address,
		},
	}

	tx = &Transaction{
		TXID:      nil,
		TXInputs:  inputs,
		TXOutPuts: outputs,
	}

	tx.SetTXID()

	return tx
}
