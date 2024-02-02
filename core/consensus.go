package core

import "fmt"

// ValidProof 方法用于验证工作量证明的有效性，接受上一个区块的工作量证明、当前计算的工作量证明和上一个区块的哈希值作为参数，返回一个布尔值表示验证结果
func (b *BlockChain) ValidProof(lastProof int, proof int, lastHash string) bool {
	// 将上一个区块的工作量证明、当前计算的工作量证明和上一个区块的哈希值拼接为字符串
	guess := fmt.Sprintf("%d%d%s", lastProof, proof, lastHash)
	// 计算拼接字符串的哈希值
	guessHash := Sha256(guess)
	// 验证哈希值的前四位是否为 "0000"，表示工作量证明有效
	return guessHash[:4] == "0000"
}

// Pow 用于执行工作量证明，接受上一个区块作为参数，返回计算得到的工作量证明数值
func (b *BlockChain) Pow(lastBlock *Block) int {
	// 获取上一个区块的工作量证明和哈希值
	lastProof := lastBlock.Proof
	lastHash := Hash(lastBlock)

	proof := 0
	// 循环直到找到有效的工作量证明
	for !b.ValidProof(lastProof, proof, lastHash) {
		proof += 1
	}

	// 返回计算得到的工作量证明数值
	return proof
}
