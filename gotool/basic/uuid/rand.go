package uuid

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandCode 指定指定位数的随机字符串
func RandCode(size int) string {
	code := ""
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < size; i++ {
		index := rand.Intn(len(chars))
		code += string(chars[index])
	}

	return code
}

// RandStr 随机字符串
// size长度
// pattern模式 1纯数字 2纯小写 4纯大写 3数字+小写 5数字+大写 6大写+小写 7全部
func RandStr(size int, pattern int) string {
	num := "0123456789"
	lowcase := "abcdefghijklmnopqrstuvwxyz"
	upcase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	code := ""
	chars := ""
	if pattern&1 == 1 {
		chars += num
	}
	if pattern&2 == 2 {
		chars += lowcase
	}
	if pattern&4 == 4 {
		chars += upcase
	}

	for i := 0; i < size; i++ {
		index := rand.Intn(len(chars))
		code += string(chars[index])
	}
	return code
}

// RandNum 获取随机数 返回 min<= x < max
func RandNum(min int64, max int64) int64 {
	if max-min <= 0 {
		return min
	}
	return rand.Int63n(max-min) + min
}
