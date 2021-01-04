package main

import "fmt"

func main() {
	//input := []int{10,9,2,5,3,7,101,18}
	//fmt.Println(lengthOfLIS(input) == 4)

	// 2, 5, 3, 4 不能识别成 2, 5 要识别成 2, 3, 4
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,4}))
	fmt.Println(lengthOfLIS([]int{4,10,4,3,8,9}))
}

// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
// 输入: [10,9,2,5,3,7,101,18]
// 输出: 4
// 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
//
// 说明:
//
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
// 你算法的时间复杂度应该为 O(n2) 。
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 注意全初始化为1
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ans := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
				ans = max(dp[i], ans)
			}
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
