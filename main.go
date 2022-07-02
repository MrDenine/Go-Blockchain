package main

import (
	//prototype "Go-Blockchain/src/prototype"
	pow "Go-Blockchain/src/proof-of-work"
)

func main() {

	bc := pow.NewBlockChain()

	bc.AddBlock("Send 1 BTC to JHON")
	bc.AddBlock("Send 2 More BTC to JHON")
	bc.GetAllBlocks()
}
