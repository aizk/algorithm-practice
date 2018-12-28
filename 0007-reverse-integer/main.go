package main

import (
	"strconv"
	"log"
	"math"
)

func main() {
	log.Println(reverse(-1010))
}

// 反转整数
// 注意：
// 	- 正负号
// 	- 最后一位 0
func reverse(x int) int {
	if -10 < x && x < 10 {
		return x
	}

	neg := false
	if x < 0 {
		x = -x
		neg = true
	}

	xStr := strconv.Itoa(x)
	xLength := len(xStr)
	reverseX := make([]rune, xLength)
	xLength--
	for _, x := range xStr {
		reverseX[xLength] = x
		xLength--
	}
	r, _ := strconv.Atoi(string(reverseX))
	if neg {
		r = -r
	}

	if r > int(math.Pow(2, 31)) {
		return 0
	}

	if r < -int(math.Pow(2, 31)) {
		return 0
	}

	return r
}

// %10 方法
func reverse2(x int) int {
	neg := 1
	if x < 0 {
		neg = -1
		x = -x
	}

	res := 0
	for x > 0 {
		// 取出 x 最后一位
		temp := x%10
		// 去除 x 最后一位
		x = x/10
		// 放入 res
		res = res * 10 + temp
	}

	res = neg * res

	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}
	return res
}