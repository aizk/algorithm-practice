// https://leetcode-cn.com/problems/coin-lcci/
package main


import "fmt"

var M = []int{1, 5, 10, 20}

func main() {
	fmt.Println(DP(100))
}

// 需要考虑 <5、<10、<20 的边缘情况，此时 sum/M[3] == 0，会少一层循环
func Force(sum int) (times int) {
	if sum < 5 {
		return sum
	}

	if sum < 10 {
		return sum
	}
	// 20 的次数
	count := 0
	for i := 0; i < sum/M[3]; i++ {
		// 10 的次数
		for j := 0; j < sum/M[2]; j++ {
			// 5 的次数
			for k := 0; k < sum/M[1]; k++ {
				// 1 的次数可以省略
				count++
				n := sum - (i * 20 + j * 10 + k * 5)
				if n >= 0 {
					times++
				}
			}
		}
	}
	fmt.Println("count: ", count)
	fmt.Println("times: ", times)
	return
}

func DP(n int) (times int) {
	moneys := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	dp[0] = 1
	for _, money := range moneys {
		for j := money; j <= n; j++ {
			dp[j] = dp[j] + dp[j - money]
		}
	}
	return dp[n]
}