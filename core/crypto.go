package core

import (
	sha2562 "crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// Hash 接收一个 *Block 类型的参数，计算该区块的哈希值并返回一个字符串
func Hash(block *Block) string {
	// 将区块结构体一转换为 JSON 格式的字符串
	blockString, _ := json.Marshal(*block)
	// 计算哈希值
	return Sha256(string(blockString))
}

// Sha256 接收一个字符串，计算该字符串的 SHA-256 哈希值并返回一个字符串
func Sha256(str string) string {
	// 创建一个 SHA-256 散列实例
	hash256 := sha2562.New()
	// 将字符串转换为字节数组，并通过 SHA-256 算法计算哈希值
	hash256.Write([]byte(str))
	// 返回计算得到的哈希值的十六进制表示
	return hex.EncodeToString(hash256.Sum(nil))
}
