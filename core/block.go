package core

// Block 表示区跨链中的一个区块
type Block struct {
	Index        int          `json:"index"`         // 区块在区块链中的位置
	Timestamp    string       `json:"timestamp"`     // 区块的时间戳，表示区块创建的时间
	Transaction  *Transaction `json:"transaction"`   // 区块的交易信息
	Proof        int          `json:"proof"`         // 工作量证明的值，用于验证新区块的合法性
	PreviousHash string       `json:"previous_hash"` // 前一个区块的哈希值，链接到上一个区块
}

// Transaction 表示区块链中的交易
type Transaction struct {
	Account   int    `json:"account"`   // 交易的金额
	Recipient string `json:"recipient"` // 交易的接收者地址
	Sender    string `json:"sender"`    // 交易的发送者地址
}
