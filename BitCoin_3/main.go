package main

func main() {

	bc := NewBlockChain("berny")
	cli := CLI{bc}
	cli.Run()
}

//测试
// ./BitCoin_2 addBlock --data "第二个区块"
