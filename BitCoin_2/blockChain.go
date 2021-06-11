package main

import (
	"log"

	"github.com/bernyw/studygo/BitCoin_2/bolt"
)

const blockChianDb = "blockChain.db"
const blockBucket = "blockBucket"

//使用数值引入区块链
type BlockChain struct {
	//使用数据库代替数组
	//key是区块的hash值，value为区块的字节流
	db *bolt.DB
	//存储最后一个区块的哈希
	tail []byte
}

//创建区块链
func NewBlockChain() *BlockChain {

	//最后一个区块的哈希,从数据库中读出来的
	var lastHash []byte

	//打开数据库
	db, err := bolt.Open(blockChianDb, 0600, nil) //数据库命名为blockChianDb，权限0600读写，无多余设置
	//defer db.Close()
	if err != nil {
		log.Panic(err)
	}

	//写数据
	db.Update(func(tx *bolt.Tx) error {
		//bucket指向，名为“blockBucket”的表
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//没有找到bucket就创建，用于存放键值对
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}

			//定义创世块
			genesisBlock := GenesisBlock()
			//block的哈希作为key，block的字节流作为value
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			//修改最后一个区块的哈希
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash

			//测试
			//blockBytes := bucket.Get(genesisBlock.Hash)
			//block := Deserialize(blockBytes)
			//fmt.Printf("block info: %v\n", block)
		} else {
			//找到bucket，已有的链，进行追加即可
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		//return nil代表整个事务操作完成，不需要回滚
		return nil
	})
	//返回刚刚操作的区块链
	return &BlockChain{
		db:   db,
		tail: lastHash,
	}
}

//定义创世块
func GenesisBlock() *Block {
	return NewBlock("创世块！\n", []byte{})

}

//添加区块到区块链
func (bc *BlockChain) AddBlock(data string) {

	//获取区块链
	db := bc.db
	//获取最后一个区块哈希
	lastHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {

		//完成区块添加
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("错误！bucket 不应为空")
		}

		//1. 创建新区块
		block := NewBlock(data, lastHash)

		//2. 添加区块到数据库中
		//hash作为key, block的字节流作为value
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)

		//3. 更新内存中的区块链
		bc.tail = block.Hash

		return nil
	})
}
