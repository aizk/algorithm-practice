package main

import "fmt"

func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		value := nums[mid]
		if value == target {
			return mid
		}

		// left in order
		// 等号是魔鬼，可能会有只有 3 个数字的情况，前面两个相等的情况
		if value >= nums[left] {
			// check target in left
			if target < value && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// check target in right
			if target > value && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
