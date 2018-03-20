package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func init() {
	GenesisBlock.Timestamp = time.Now().UnixNano()
}

// Multi-party transaction can be simplified into multiple 1:1 transactions.
type Transaction struct {
	Owes   string  `json:"owes,omitempty"`
	To     string  `json:"to,omitempty"`
	Amount float32 `json:"amount,omitempty"`
}

type Transactions [5]Transaction

// Block is a block in the blockchain which contains number of transactions.
// We limit that to 5 transactions per block.
type Block struct {
	Index     int
	Nonce     int
	Timestamp int64
	Txns      Transactions `json:"txns,omitempty"`
	Hash      string
	PrevHash  string
}

// New takes the block index, previous block hash and group of 5 transactions and returns a block.
func New(index int, prevHash string, txns Transactions) Block {
	b := Block{
		Index:     index,
		Nonce:     0, // TODO: add proof of work later.
		Timestamp: time.Now().UnixNano(),
		Txns:      txns,
		PrevHash:  prevHash,
	}

	// Calculate the hash of the block with all the fields except the current hash included.
	b.Hash = b.GetHash()

	return b
}

const (
	genesisPrevHash = "0000000000000000000000000000000000000000000000000000000000000000"
)

var GenesisBlock = Block{
	Index: 0,
	Nonce: 0,
	// Timestamp set during initialization.
	Txns: Transactions{},
	PrevHash: genesisPrevHash,
}

// Blockchain is a slice of blocks.
type Blockchain []Block

// NewBlockchain returns a new Blockchain initialized with the Genesis block.
func NewBlockchain() Blockchain {
	GenesisBlock.Hash = GenesisBlock.GetHash()
	return Blockchain{GenesisBlock}
}

func (tx Transactions) String() string {
	b, err := json.Marshal(tx)
	if err != nil {
		log.Fatalf("error marshaling transactions: %s", err)
	}

	return string(b)
}

func (b Block) String() string {
	// We can't include the hash field itself in the hash calculation.
	// It'd be a chicken and the egg problem.
	return fmt.Sprintf("%d:%d:%d:%s:%s", b.Index, b.Nonce, b.Timestamp, b.Txns.String(), b.PrevHash)
}

func (b Block) PrettyPrint() {
	 fmt.Printf("Index: %10d\nNonce: %10d\nTimestamp: %10d\nTransactions: %10s\nPrevHash: %10s\nHash: %10s\n", b.Index, b.Nonce, b.Timestamp, b.Txns.String(), b.PrevHash, b.Hash)
}

func (b Block) GetHash() string {
	h := sha256.New()
	h.Write([]byte(b.String()))
	return hex.EncodeToString(h.Sum(nil))
}

func (bc *Blockchain) Append(b Block) bool {
	if b.PrevHash == bc.LastBlock().Hash && b.Index == bc.LastBlock().Index + 1 && b.Hash == b.GetHash(){
		*bc = append(*bc, b)
		return true
	}

	log.Printf("Offending block:")
	b.PrettyPrint()
	return false
}

func (bc Blockchain) LastBlock() Block {
	return bc[len(bc)-1]
}


func (bc Blockchain) PrettyPrint() {
	for _, b := range bc {
		b.PrettyPrint()
	}
}