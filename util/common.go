package util

import (
	"math/rand"
	"strconv"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// UintToStr convert uint to string
func UintToStr(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}

func StrToUint(s string) uint64 {
	n, _ := strconv.ParseUint(s, 10, 64)
	return n
}
