package main

import (
	"fmt"
	"os"
)

const Usage = `
	./blockchain addBlock "xxxx"	添加数据到区块链
	./blockchain printChain			打印区块链
	./blockchain getBalance	地址		获取该地址的余额
`

type CLI struct {
	bc *BlockChain
}

//给CLI提供一个方法，进行命令解析，从而执行调度
func (cli *CLI) Run() {
	cmds := os.Args

	if len(cmds) < 2 {
		fmt.Printf(Usage)
		os.Exit(1)
	}

	switch cmds[1] {
	case "addBlock":
		if len(cmds) != 3 {
			fmt.Printf(Usage)
			os.Exit(1)
		}

		fmt.Printf("添加区块命令被调用, 数据:%s\n", cmds[2])
		//data := []byte(cmds[2]) //TODO
		//cli.AddBlock(data) //TODO
	case "printChain":
		fmt.Printf("打印区块链命令被调用\n")
		cli.PrintChain()
	case "getBalance":
		fmt.Printf("获取余额命令被调用\n")
		cli.bc.GetBalance([]byte(cmds[2]))
		
	default:
		fmt.Println("无效的命令，请检查")
		fmt.Printf(Usage)
	}
}
