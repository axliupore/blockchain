package core

import (
	"encoding/json"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"net/http"
	"time"
)

// Chain 表示区块链中的链
type Chain struct {
	Chain  []*Block `json:"chain"`  // 区块链中的所有区块
	Length int      `json:"length"` // 区块链的长度，即区块的数量
}

// BlockChain 表示一个完整的区跨链
type BlockChain struct {
	CurrentTransaction []*Transaction // 当前待处理的交易列表
	Chain              []*Block       // 区块链中的所有区块
	Nodes              mapset.Set     // 区块链中网络中的节点集合
}

// New 创建一个新的区块链
func New() *BlockChain {
	b := &BlockChain{}   // 创建一个新的区块链
	b.NewBlock(100, "1") // 生成创世区块
	return b
}

// NewBlock 创建新的区块
func (b *BlockChain) NewBlock(proof int, previousHash string) *Block {
	block := &Block{}
	// 设置区块的索引
	block.Index = len(b.Chain) + 1

	// 设置区块的时间戳
	block.Timestamp = time.Now().String()

	// 设置区块的工作量证明和前一个区块的哈希值
	block.Proof = proof
	if previousHash != "" {
		block.PreviousHash = previousHash
	} else {
		block.PreviousHash = Hash(b.Chain[len(b.Chain)-1])
	}

	// 如果有待处理的交易，则将第一笔交易添加到区块中
	if len(b.CurrentTransaction) != 0 {
		block.Transaction = b.CurrentTransaction[0]
	}
	b.CurrentTransaction = []*Transaction{}

	// 将新的区块添加到区块链中
	b.Chain = append(b.Chain, block)
	return block
}

// LastBlock 返回区块链中的最后一个区块
func (b *BlockChain) LastBlock() *Block {
	return b.Chain[len(b.Chain)-1]
}

// ValidChain 验证一个区块链的合法性
func (b *BlockChain) ValidChain(chain []*Block) bool {
	lastBlock := b.Chain[1] // 获取链中的第一个区块作为起点
	currentIndex := 2       // 从第二个区块开始遍历

	for currentIndex < len(chain) { // 遍历整个链
		block := chain[currentIndex]
		fmt.Println(lastBlock, block)

		// 检查前一个区块的哈希值是否等于当前区块中记录的前一个区块的哈希值
		if block.PreviousHash != Hash(lastBlock) {
			return false
		}

		// 检查工作量证明的合法性
		if !b.ValidProof(lastBlock.Proof, block.Proof, lastBlock.PreviousHash) {
			return false
		}

		lastBlock = block // 更新当前区块为上一个区块
		currentIndex += 1 // 遍历下一个区块
	}
	return true
}

// NewTransaction 用于创建一个新的交易
func (b *BlockChain) NewTransaction(sender string, recipient string, amount int) int {
	transaction := &Transaction{}

	// 设置交易的发送者、接收者和金额
	transaction.Sender = sender
	transaction.Account = amount
	transaction.Recipient = recipient

	// 将交易添加到当前待处理的交易列表中
	b.CurrentTransaction = append(b.CurrentTransaction, transaction)
	return b.LastBlock().Index + 1
}

// RegisterNode 注册新的节点到区块链网络
func (b *BlockChain) RegisterNode(host string) {
	if b.Nodes == nil { // 如果节点集合为空
		b.Nodes = mapset.NewSet() // 创建新的节点集合
	}
	b.Nodes.Add(host) // 将新节点添加到集合中
}

// ResolveConflicts 用于解决不同节点之间的区块链冲突
func (b *BlockChain) ResolveConflicts() bool {
	maxLength := len(b.Chain) // 当前链的长度

	var newChain []*Block

	for node := range b.Nodes.Iter() { // 遍历网络中的每个节点
		resp, err := http.Get("http://" + node.(string) + "/chain")
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		chain := &Chain{}

		// 解码响应的区块链
		err = json.NewDecoder(resp.Body).Decode(chain)
		if err != nil {
			return false
		}
		// 如果从其他节点获得的区块链比当前的长且合法，则更新区块链
		if chain.Length > maxLength && b.ValidChain(chain.Chain) {
			maxLength = chain.Length
			newChain = make([]*Block, len(chain.Chain))
			copy(newChain, chain.Chain)
		}
	}

	// 如果找到了合法的更长的区块链，则更新当前的区块链
	if newChain != nil {
		b.Chain = make([]*Block, len(newChain))
		copy(b.Chain, newChain)
		return true
	}
	return false
}
