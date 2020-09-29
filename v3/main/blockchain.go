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

func NewBlockChain(address string) *BlockChain {
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
			//创世块中只有一个挖矿交易，只有coinBase
			coinBase := NewCoinbaseTx(address)
			genesisBlock := NewBlock([]*Transaction{coinBase}, nil)

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
		} else {
			blockChain.lastHash = b.Get([]byte(lastHashKey))
			blockChain.db = db
		}

		return err
	})

	return blockChain
}

//添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	bc.db.Update(func(tx *bolt.Tx) error {
		//所有的操作都在这里
		b := tx.Bucket([]byte(blockBucketName))

		if b != nil {
			//bucket已经创建完成
			block := NewBlock(txs, bc.lastHash)

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

func (bc *BlockChain) FindMyUtxos(address []byte) (ret []TXOutput) {
	fmt.Printf("FindMyUtxos\n")
	var inputs []TXInput

	//1.遍历账本
	it := bc.NewIterator()
	for  {
		block := it.Next()
		//2.遍历交易
		for _, tx := range block.Transactions {
			//遍历交易输入：inputs
			for _, input := range tx.TXInputs {
				if  string(address) == input.Address {
					inputs = append(inputs, input)
				}
			}

			//3.遍历output
			for i, output := range tx.TXOutPuts {
				//4.找到属于我的所有output
				if  string(address) == output.Address {
					fmt.Printf("找到了属于:%s的output，i: %d\n", address, i)
					ret = append(ret, output)
				}
			}
		}

		if block.PrevBlockHash == nil {
			fmt.Println("遍历区块链结束")
			break
		}
	}

	return
}

func (bc *BlockChain) GetBalance(address []byte) (ret float64) {
	utxos := bc.FindMyUtxos(address)

	for _, utxo := range utxos {
		ret += utxo.Value
	}

	fmt.Printf("%s:的余额为：%f\n", address, ret)
	return
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
