package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义proofOfWork结构
type ProofOfWork struct {
	//a. block
	block *Block
	//b. 目标值
	target *big.Int
}

//提供创建POW函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//难度值为16进制
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"

	//引入辅助变量tmpInt，将targetStr抓换为big.int
	tmpInt := big.Int{}
	//将难度值赋值给big.int，指定16进制的格式
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow
}

//计算函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//拼装数据
	var nonce uint64
	block := pow.block
	//存储计算结果
	var hash [32]byte

	fmt.Println("开始挖矿...")
	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			//block.Data,
			//只对区块头做哈希值，区块体通过MerkelRoot产生影响
		}
		//将二维的切片数组链接起来，无分隔符
		blockInfo := bytes.Join(tmp, []byte{})
		//做hash运算
		hash = sha256.Sum256(blockInfo)
		//与pow中的target进行比较
		tmpInt := big.Int{}
		//将得到的hash数组转换成一个big.int
		tmpInt.SetBytes(hash[:])
		//与目标值比较
		if tmpInt.Cmp(pow.target) == -1 {
			//当 tmpInt <  pow.target 找到了
			fmt.Printf("挖矿成功！ hash: %x, nonce: %d\n", hash, nonce)
			return hash[:], nonce
		} else {
			//没到到 nonce加1
			nonce++
		}

	}
}
