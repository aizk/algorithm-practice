package main

import "log"

func main() {
	log.Println(numJewelsInStones("aA", "aAA"))
}

// J 里面的字符是珠宝
// S 是一堆石头
// 判断 S 中珠宝的数量
func numJewelsInStones(J string, S string) int {
	if len(J) == 0 || len(S) == 0 {
		return 0
	}
	jewels := make(map[rune]interface{})
	for _, j := range J {
		jewels[j] = j
	}

	count := 0
	for _, s := range S {
		_, ok := jewels[s]
		if ok {
			count++
		}
	}
	return count
}

func numJewelsInStones2(J string, S string) int {
	var diamondMap = map[byte]bool{}
	for i := 0; i < len(J); i++ {
		diamondMap[J[i]] = true
	}

	cnt := 0
	for i := 0; i < len(S); i++ {
		if diamondMap[S[i]] == true {
			cnt++
		}
	}
	return cnt
}