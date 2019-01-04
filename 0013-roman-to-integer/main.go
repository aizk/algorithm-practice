package main

import "log"

func main() {
	log.Println(romanToInt("III"))
	log.Println(romanToInt("VI"))
}

var rule = map[rune]int{
	'I':   1,
	'V':   5,
	'X':   10,
	'L':   50,
	'C':   100,
	'D':   500,
	'M':   1000,
}

// 罗马数字都是从大到小书写，也就是说不会出现 VX 这种 5 在 10 左边的情况，因此可以实现
// 逐个分析字符串的字符，根据规则递加得到结果
// 关键是抓住规律，并抽象出规律
// 先把每一位都加上，然后把要减去的减掉
func romanToInt(s string) int {
	rsp := 0
	for i, n := range s {
		// 字符串位置不到最后一位 且 上一位的值小于当前位的值
		if i != 0 && rule[rune(s[i-1])] < rule[n] {
			// 减去上一位的值、假设是 XC，X 的时候加 10，C 的时候加 100，算下来多加了 20，所以要减去两倍的量
			rsp -= 2 * rule[rune(s[i-1])]
		}
		// 加上当前位的值
		rsp += rule[n]
	}
	return rsp
}

