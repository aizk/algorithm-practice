package main

import "fmt"

func main() {
	TestIsDeformation("xyz", "xyz", true)
	TestIsDeformation("xyz", "xyzz", false)
	TestIsDeformation("", "", false)
}

func TestIsDeformation(x, y string, expected bool) {
	r := IsDeformation(x, y)
	if r != expected {
		fmt.Printf("input: %s and %s | result: %t | expected: %t", x, y, r, expected)
	}
}

// 互为变形词
// 长度不相等，直接 GG
// 长度相等，减少到 -1 也 GG
func IsDeformation(x, y string) (bool) {
	if len(x) != len(y) {
		return false
	}

	if len(x) == 0 || len(y) == 0 {
		return false
	}

	m := map[rune]int{}
	// total x byte count
	for _, b := range x {
		m[b]++
	}

	for _, b := range y {
		m[b]--
		if m[b] < 0 {
			return false
		}
	}
	return true
}