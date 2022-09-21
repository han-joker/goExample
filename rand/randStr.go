package rand

import (
	"math/rand"
	"strings"
)

func RandStrMaskSB(l int) string {

	// 预设字符集合
	const digits = "0123456789"
	const letters = "abcdefghijklmnopqrstuvwxyz"

	const bytes = digits + letters
	bytesL := len(bytes)
	// 获取字符集合需要的二进制长度
	const (
		// 二进制长度
		letterIdxBits = 5
		// 形成掩码
		letterIdxMask = 1<<letterIdxBits - 1
		// 63 位可以使用的最大组次数
		letterIdxMax = 63 / letterIdxBits
	)

	// 利用string builder构建结果缓冲
	sb := strings.Builder{}
	sb.Grow(l)
	// 循环生成随机数
	// i 索引
	// cache 随机数缓存
	// remain 随机数还可以用几次
	for i, cache, remain := l-1, rand.Int63(), letterIdxMax; i >= 0; {
		// 随机缓存不足，重新生成
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		// 利用掩码生成随机索引，有效索引为小于字符集合长度
		if idx := int(cache & letterIdxMask); idx < bytesL {
			sb.WriteByte(bytes[idx])
			i--
		}
		// 利用下一组随机数位
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func RandStrSimple(l int) string {
	// 预设字符集合
	const digits = "0123456789"
	const letters = "abcdefghijklmnopqrstuvwxyz"
	const bytes = digits + letters

	bytesL := len(bytes)

	sb := strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteByte(bytes[rand.Intn(bytesL)])
	}
	return sb.String()
}
