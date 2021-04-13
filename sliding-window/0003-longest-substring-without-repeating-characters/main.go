package main

func main() {

}

func lengthOfLongestSubstring(s string) (res int) {
	window := make(map[uint8]int)
	left, right := 0, 0

	for right < len(s) {
		c := s[right]
		right++
		_, ok := window[c]
		if ok {
			window[c]++
		} else {
			window[c] = 1
		}
		for window[c] > 1 {
			// 收缩
			d := s[left]
			left++
			window[d]--
		}
		// 更新答案
		if res < right - left {
			res = right - left
		}
	}

	return
}