package main

//使用数值引入区块链
type BlockChain struct {
	//定义一个区块类型的数组
	blocks []*Block
}

//创建区块链
func NewBlockChain() *BlockChain {
	//定义创世块
	genesisBlock := GenesisBlock()

	//返回一个区块链
	return &BlockChain{blocks: []*Block{genesisBlock}}
}

//定义创世块
func GenesisBlock() *Block {
	//date“创世块！”，前一个区块哈希为空
	return NewBlock("创世块！", []byte{})

}

//添加区块到区块链
func (bc *BlockChain) AddBlock(data string) {
	//获取前区块哈希
	//获取前一个区块位置
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash
	//	a. 创建新的区块
	block := NewBlock(data, prevHash)
	//	b. 添加区块到区块链数组中
	bc.blocks = append(bc.blocks, block)
}
