package _567_permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[rune]int, len(s1))
	window := make(map[rune]int, len(s1))
	for _, c := range s1 {
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
	sr := []rune(s2)

	for right < len(sr) {
		c := sr[right]
		right++
		_, ok := need[c]
		if ok {
			_, ok = window[c]
			if ok {
				window[c]++
			} else {
				window[c] = 1
			}

			if window[c] == need[c] {
				valid++
			}
		}

		for right - left >= len(s1) {
			// 找到符合的排列子串
			if valid == len(need) {
				return true
			}
			d := sr[left]
			left++
			_, ok = need[d]
			if ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}
