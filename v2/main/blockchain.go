package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

//创建区块链，使用bolt数据库
type BlockChain struct {
	db       *bolt.DB
	lastHash []byte
}

const blockChainName = "blockChain.db"
const blockBucketName = "blockBucket"
const lastHashKey = "lastHashKey"

func NewBlockChain() *BlockChain {
	//在创建的时候添加一个区块：创世块
	blockChain := &BlockChain{}
	db, err := bolt.Open(blockChainName, 0600, nil)

	if err != nil {
		panic(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		//所有的操作都在这里
		b := tx.Bucket([]byte(blockBucketName))

		if b == nil {
			b, err = tx.CreateBucket([]byte(blockBucketName))
			if err != nil {
				panic(err)
			}

			//bucket已经创建完成
			genesisBlock := NewBlock([]byte("genius block"), nil)

			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				fmt.Println("写入数据失败: name1->lily")
			}

			err = b.Put([]byte(lastHashKey), genesisBlock.Hash)

			if err != nil {
				fmt.Println("写入数据失败: name2->Bob")
			}

			if err == nil {
				blockChain.lastHash = genesisBlock.Hash
				blockChain.db = db
			}

			//为了测试，我们把写入的数据读取出来，如果没有问题再注释掉
			//blockInfo := b.Get(blockChain.lastHash)
			//block := Deserialize(blockInfo)
			//fmt.Printf("解码后的block.Data:%s\n", block.Data)
		} else {
			blockChain.lastHash = b.Get([]byte(lastHashKey))
			blockChain.db = db
		}

		return err
	})

	return blockChain
}

//添加区块
func (bc *BlockChain) AddBlock(data []byte) {
	bc.db.Update(func(tx *bolt.Tx) error {
		//所有的操作都在这里
		b := tx.Bucket([]byte(blockBucketName))

		if b != nil {
			//bucket已经创建完成
			block := NewBlock(data, bc.lastHash)

			err := b.Put(block.Hash, block.Serialize())
			if err != nil {
				fmt.Println("写入数据失败:")
			} else {
				err = b.Put([]byte(lastHashKey), block.Hash)

				if err != nil {
					fmt.Println("写入数据失败:")
				} else {
					bc.lastHash = block.Hash
				}
			}
		} else {
			panic("bucket is not exists")
		}

		return nil
	})
}

//定义一个区块链的迭代器，包含db和current
type BlockChainIterator struct {
	db      *bolt.DB
	current []byte //当前所指向的区块的哈希值
}

//创建迭代器，使用bc进行初始化
func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		db:      bc.db,
		current: bc.lastHash,
	}
}

func (it *BlockChainIterator) Next() (block *Block) {
	//访问数据库，拿到最后一个区块
	it.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucketName))

		if b == nil {
			panic("bucket not exists")
		}

		blockHash := b.Get(it.current)
		block = Deserialize(blockHash)
		it.current = block.PrevBlockHash

		return nil
	})

	return
}
