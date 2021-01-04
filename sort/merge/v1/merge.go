package main

import "fmt"

func main() {
	data := []int{5, 4, 3, 2, 1}
	result := sort(data)
	fmt.Println(result)
}

// 这个方法空间复杂度不低（开辟很多新的数组），但是逻辑清晰
// 在 v2 中优化为只使用一个额外数组
func sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	midd := len(arr) / 2
	return merge(sort(arr[:midd]), sort(arr[midd:]))
}

func merge(l, r []int) (result []int) {
	for len(l) != 0 && len(r) != 0 {
		if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	if len(l) != 0 {
		result = append(result, l...)
	} else if len(r) != 0 {
		result = append(result, r...)
	}
	return
}