package pkg

import (
	"crypto/sha256"
	"fmt"
	"io"
	"time"
)

type Bloc struct {
	Index     int
	Timestamp time.Time
	Data      string
	PrevHash  string
	Hash      string
}

func NewBlock(data string, prevBloc *Bloc) *Bloc {
	bloc := &Bloc{
		Timestamp: time.Now(),
		Data:      data,
	}

	if prevBloc == nil {
		bloc.Index = 0
		bloc.PrevHash = ""
		bloc.Hash = bloc.CalculateHash()
		return bloc
	}

	bloc.Index = prevBloc.Index + 1
	bloc.PrevHash = prevBloc.Hash
	bloc.Hash = bloc.CalculateHash()
	return bloc
}

func (bloc *Bloc) CalculateHash() string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%d%s%s%s", bloc.Index, bloc.Timestamp, bloc.Data, bloc.PrevHash)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

type BlockChain struct {
	blockChain []*Bloc
}

func NewBlockChain() *BlockChain {
	return &BlockChain{blockChain: []*Bloc{NewBlock("Genesis Bloc", nil)}}
}

func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blockChain[len(bc.blockChain)-1]
	newblock := NewBlock(data, lastBlock)
	bc.blockChain = append(bc.blockChain, newblock)
}

func (bc *BlockChain) Print(w io.Writer) {
	for _, block := range bc.blockChain {
		fmt.Fprintf(w, "Index: %d\n", block.Index)
		fmt.Fprintf(w, "Timestamp: %s\n", block.Timestamp)
		fmt.Fprintf(w, "Data: %s\n", block.Data)
		fmt.Fprintf(w, "PrevHash: %s\n", block.PrevHash)
		fmt.Fprintf(w, "Hash: %s\n", block.Hash)
		fmt.Fprintln(w)
	}
}
