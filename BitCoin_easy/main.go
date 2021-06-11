package main

import (
	"fmt"
)

func main() {
	//创建区块链，初始化创世区块
	bc := NewBlockChain()
	bc.AddBlock("第二个区块")
	//遍历区块链
	for i, block := range bc.blocks {
		fmt.Printf("============当前区块高度%d==============\n", i)
		fmt.Printf("前区块哈希: %x\n", block.PrevHash)
		fmt.Printf("区块哈希: %x\n", block.Hash)
		fmt.Printf("区块数据: %s\n", block.Data)
	}

}
