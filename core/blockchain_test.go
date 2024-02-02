package core

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewBlock(t *testing.T) {
	b := New()
	spew.Dump(b.Chain[0])
	block := b.NewBlock(1, "")
	spew.Dump(block)
	block = b.NewBlock(2, "")
	spew.Dump(block)
}

func TestValidChain(t *testing.T) {
	b := New()
	var blocks []*Block
	blocks = append(blocks, b.Chain[0])
	for i := 0; i < 4; i++ {
		lastBlock := b.LastBlock()
		block := b.NewBlock(b.Pow(lastBlock), Hash(lastBlock))
		blocks = append(blocks, block)
	}
	if b.ValidChain(blocks) {
		for _, block := range blocks {
			spew.Dump(block)
		}
	}
}
