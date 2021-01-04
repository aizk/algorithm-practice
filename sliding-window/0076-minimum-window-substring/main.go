package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC") == "BANC")
}

func minWindow(s string, t string) string {
	need := make(map[rune]int, len(t))
	window := make(map[rune]int, len(t))
	for _, c := range t {
		_, ok := need[c]
		if ok {
			need[c]++
		} else {
			need[c] = 1
		}
	}

	var left, right int
	// 表示窗口中满足 need 条件的字符个数
	valid := 0
	// 记录最小覆盖子串的起始索引和长度
	start := 0
	lenght := math.MaxInt64

	sr := []rune(s)

	for right < len(sr) {
		// c 是将移入窗口的字符
		c := sr[right]
		// 右移窗口
		right++
		// c 是不是需要的字符
		_, ok := need[c]
		if ok {
			_, ok = window[c]
			if ok {
				window[c]++
			} else {
				window[c] = 1
			}
			// window 满足 need 条件
			if window[c] == need[c] {
				valid++
			}
		}

		// 是否都满足条件，开始收缩左侧
		// 所有需要的字符都满足
		for valid == len(need) {
			// 更新最小覆盖子串
			if right - left < lenght {
				start = left
				lenght = right - left
			}
			// 移除一个字符
			d := sr[left]
			left++
			_, ok := need[d]
			if ok {
				// 如果相等，收缩后就不满足了
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if lenght == math.MaxInt64 {
		return ""
	} else {
		return string(sr[start: start + lenght])
	}
}