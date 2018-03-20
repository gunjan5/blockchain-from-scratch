package main

import (
	"fmt"

	"github.com/gunjan5/blockchain-from-scratch/block"
)

func main() {

	// Initialize the blockchain with the Genesis block.
	bc := block.NewBlockchain()

	lastSuccessfulBlock := block.GenesisBlock

	lastSuccessfulBlock, _ = addBlock(&bc, block.Transactions{
		block.Transaction{
			Owes:   "Gunjan",
			To:     "Bob",
			Amount: 4.20,
		},
		block.Transaction{
			Owes:   "Elon",
			To:     "Gunjan",
			Amount: 3.14,
		},
		block.Transaction{
			Owes:   "Lannister",
			To:     "No-one",
			Amount: 1000.0,
		},
	}, lastSuccessfulBlock)

	lastSuccessfulBlock, _ = addBlock(&bc, block.Transactions{
		block.Transaction{
			Owes:   "Happy",
			To:     "Lemon",
			Amount: 4.5,
		},
	}, lastSuccessfulBlock)

	fmt.Println("-----------------\nBlockchain:\n-----------------")
	bc.PrettyPrint()

}

func addBlock(bc *block.Blockchain, txns block.Transactions, lsb block.Block) (block.Block, bool) {
	b := block.New(lsb.Index+1, lsb.Hash, txns)
	if bc.Append(b) {
		lsb = b
		return lsb, true
	}

	return lsb, false
}
