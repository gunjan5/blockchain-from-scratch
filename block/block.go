package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Multi-party transaction can be simplified into multiple 1:1 transactions.
type Transaction struct {
	Owes   string
	To     string
	Amount float32
}

type Transactions [5]Transaction

// Block is a block in the blockchain which contains number of transactions.
// We limit that to 5 transactions per block.
type Block struct {
	Index     int
	Nonce     int
	Timestamp int64
	Txns      Transactions
	Hash      string
	PrevHash  string
}

// Blockchain is a slice of blocks.
type Blockchain []Block

// New takes the block index, previous block hash and group of 5 transactions and returns a block.
func New(index int, prevHash string, txns Transactions) *Block {
	b := Block{
		Index:     index,
		Nonce:     0, // add proof of work later.
		Timestamp: time.Now().Unix(),
		Txns:      txns,
		PrevHash:  prevHash,
	}

	// Calculate the hash of the block with all the fields except the current hash included.
	b.Hash = b.GetHash()

	return &b
}

func (tx Transactions) String() string {
	b, err := json.Marshal(tx)
	if err != nil {
		log.Fatalf("error marshaling transactions: %s", err)
	}

	return string(b)
}

func (b Block) String() string {
	return fmt.Sprintf("%d:%d:%d:%s:%s:%s", b.Index, b.Nonce, b.Timestamp, b.Txns.String(), b.PrevHash, b.Hash)
}

func (b Block) PrettyPrint() string {
	return fmt.Sprintf("Index: %10d\nNonce: %10d\nTimestamp: %10d\nTransactions: %10s\nPrevHash: %10s\nHash: %10s\n", b.Index, b.Nonce, b.Timestamp, b.Txns.String(), b.Hash, b.PrevHash)
}

func (b Block) GetHash() string {
	h := sha256.New()
	h.Write([]byte(b.String()))
	return hex.EncodeToString(h.Sum(nil))
}
