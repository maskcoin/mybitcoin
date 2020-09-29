package main

//创建区块链，使用Block数组模拟
type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	//在创建的时候添加一个区块：创世块
	blockChain := &BlockChain{}
	block := NewBlock([]byte("genius block"), nil)
	blockChain.Blocks = append(blockChain.Blocks, block)
	return blockChain
}

//添加区块
func (bc *BlockChain) AddBlock(data []byte) {
	//1.创建一个区块
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data, lastBlock.Hash)
	//2.添加到bc.Blocks中
	bc.Blocks = append(bc.Blocks, block)
}
