package core

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	// 创建一个区块
	b := &Block{
		Index:     1,
		Timestamp: "2024-2-1",
		Transaction: &Transaction{
			Account:   100,
			Recipient: "localhost",
			Sender:    "127.0.0.1",
		},
		Proof:        10,
		PreviousHash: "1",
	}
	// 输出 SHA-256 哈希值
	fmt.Println(Hash(b))
}
