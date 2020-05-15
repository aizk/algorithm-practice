package main

import "log"

func main() {
	log.Println(twoSum([]int{3, 2, 4}, 6))
}

// 找到数组中两个数的和为 target
func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}
	for index, num := range nums {
		numsMap[num] = index
	}

	for index, num := range nums {
		find := target - num
		idx, ok := numsMap[find]
		if idx == index {
			continue
		}
		if ok {
			return []int{index, idx}
		}
	}
	return nil
}
