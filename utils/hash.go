package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// 传入字符串，返回其hash值
func Hash(inputString string) (hashString string) {
	// 创建 SHA-256 哈希对象
	hash := sha256.New()

	// 将字符串转换为字节数组并计算哈希值
	hash.Write([]byte(inputString))
	hashValue := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString = hex.EncodeToString(hashValue)
	return
}
