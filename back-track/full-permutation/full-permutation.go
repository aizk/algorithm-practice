package main

import "fmt"

var result [][]int

func main() {
	nums := []int{1, 2, 3}

	backtrack(nums, &[]int{})

	fmt.Println(result)
}

func backtrack(nums []int, track *[]int) {
	if len(nums) == len(*track) {
		result = append(result, append([]int{}, *track...))
		return
	}

	for _, num := range nums {
		// 排除已有选择
		choosed := false
		for _, v := range *track {
			if num == v {
				choosed = true
				break
			}
		}
		if choosed {
			continue
		}

		// 选择
		*track = append(*track, num)
		// 进入下一层
		backtrack(nums, track)
		// 取消选择 remove last
		*track = (*track)[:len(*track) - 1]
	}
}
