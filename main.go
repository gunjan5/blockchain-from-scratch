package main

import (
	"fmt"

	"github.com/gunjan5/blockchain-from-scratch/block"
)

const (
	genesisPrevHash = "0000000000000000000000000000000000000000000000000000000000000000"
)

func main() {

	txns := block.Transactions{
		block.Transaction{
			Owes: "Gunjan",
			To: "Bob",
			Amount: 4.20,
		},
		block.Transaction{
			Owes: "Elon",
			To: "Gunjan",
			Amount: 3.14,
		},
		block.Transaction{
			Owes: "Lannister",
			To: "No-one",
			Amount: 1000.0,
		},
	}

	b := block.New(0, genesisPrevHash, txns)

	fmt.Printf("Genesis block:\n%s\n", b.PrettyPrint())
}
