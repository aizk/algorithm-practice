package _416

import "sync"

func main() {

}

func CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % 2 == 1 {
		return false
	}
	sum = sum / 2
	l := len(nums)

	dp := make([][]bool, l + 1)
	for i := 0; i <= l; i++ {
		ele := make([]bool, sum + 1)
		// base case 空背包表示满的
		ele[0]  = true
		dp[i] = ele
	}

	for i := 1; i <= l; i++ {
		for j := 1; j <= sum; j++ {
			if j - nums[i-1] < 0 {
				// 背包容量不足，不能装入第 i 个物品
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = hasTrue(dp[i-1][j], dp[i-1][j - nums[i-1]])
			}

		}
	}

	return dp[l][sum]
}

// 压缩掉了物品个数的状态
// 速度比二维快一倍
func 状态压缩(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % 2 == 1 {
		return false
	}
	sum = sum / 2
	n := len(nums)

	dp := make([]bool, sum + 1)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := sum; j >= 0; j++ {
			if j - nums[i] >= 0 {
				dp[j] = hasTrue(dp[j], dp[j - nums[i]])
			}
		}
	}

	return dp[sum]
}

func hasTrue(x, y bool) bool {
	if x || y {
		return true
	}
	return false
}