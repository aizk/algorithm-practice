package main

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
//
// 示例 1:
//
// 输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。
//
// 示例 2:
//
// 输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
//
// 说明: 你可以假设 n 不小于 2 且不大于 58。
// Related Topics 数学 动态规划

func main() {
	println(integerBreak(10))
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// 参考：https://mp.weixin.qq.com/s/ixnmuaj4lZJWj5jMMjnpWA
// 类似剪绳子
// 暴力递归
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	res := 0
	for i := 1; i < n; i++ {
		// 把所有结果试一遍
		// f(n) = max(1 * fn(n - 1), 2 * f(n - 2), ..., (n - 1) * f(1))
		// 如果拆分一个 1：判断是 1 * f(n - 1) 大，还是只拆分两个数大 1 * (n - 1)，因为：
		// f(n - 1) 是子问题，可能会拆成另外两个数相乘得到子问题最大！
		res = max(res, max(i * integerBreak(n - i), i * (n - i)))
	}
	return res
}




