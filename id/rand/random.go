package rand

import (
	"math/rand"
	"time"
)

var (
	defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	defaultNums    = []rune("0123456789")
)

// RandomNum
// @Description: 随机数字
// @param: n
// @return string
func RandomNum(n int) string {
	return random(n, defaultNums)
}

// RandomStr
// @Description: 随机字符串
// @param: n
// @return string
func RandomStr(n int) string {
	return random(n, defaultLetters)
}

// random
// @Description: 随机字符串
// @param: n 随机字符串长度
// @param: allowedChars 允许的字符
// @return string
func random(n int, allowedChars []rune) string {
	numSize := len(allowedChars)
	if numSize == 0 {
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	b := make([]rune, n)
	for i := range b {
		ind := r.Intn(numSize)
		b[i] = allowedChars[ind]
	}
	return string(b)
}
