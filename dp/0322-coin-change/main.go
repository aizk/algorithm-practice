package main

import "fmt"

func main() {
	fmt.Println(coinChange([]int{1, 3, 5, 7}, 100))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j - coin] + 1)
		}
	}

	if dp[amount] != amount + 1 {
		return dp[amount]
	} else {
		return -1
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}