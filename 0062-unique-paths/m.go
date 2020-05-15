package main

import "fmt"

func main() {
	uniquePaths(7, 3)
}

// 7 x 3
func uniquePaths(m int, n int) int {
	var dp [][]int
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, m))
		// 下面这样初始化不行
		//dp = append(dp, []int{})
	}

	// 第一行初始化
	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}
	// 第一列初始化
	for j := 0; j < n; j++ {
		dp[j][0] = 1
	}

	for _, row := range dp {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println("")
	}

	// 开始递推
	// i 、 j 要和初始化的方式对应，这里初始化为行列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}

	for _, row := range dp {
		for _, col := range row {
			fmt.Print(col, "|")
		}
		fmt.Println("")
	}

	return dp[n - 1][m -1]
}
