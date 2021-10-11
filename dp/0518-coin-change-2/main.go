package main

import "fmt"

func main() {
	fmt.Println(change(1000, []int{1, 2, 5})) // 50401
}


func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j] + dp[j - coin]
		}
	}
	return dp[amount]
}