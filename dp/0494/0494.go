package _494

var Count int

func Recursive(nums []int, S int) {
	if len(nums) == 0 {
		if S == 0 {
			Count++
		}
		return
	}

	num := nums[len(nums) - 1]
	nums = nums[:len(nums) - 1]

	Recursive(nums, S-num)
	Recursive(nums, S+num)
}

func recursive(nums []int, S int, count *int)  {
	if len(nums) == 0 {
		if S == 0 {
			*count++
		}
		return
	}

	num := nums[len(nums) - 1]
	nums = nums[:len(nums) - 1]

	recursive(nums, S-num, count)
	recursive(nums, S+num, count)
}

// 暴力回溯提交会超出时间限制
func FindTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	chosen := make([]int, len(nums))
	helper(&nums, &chosen, S, &count)
	return count
}

func helper(nums *[]int, chosen *[]int, S int, count *int) {
	if len(*nums) == 0 {
		// 判断是否等于 S
		sum := 0
		for _, i := range *chosen {
			sum += i
		}
		if sum == S {
			*count++
		}
		return
	}

	for _, symbol := range []int{1, -1} {
		v := (*nums)[len(*nums) - 1]
		c := symbol * v
		*nums = (*nums)[:len(*nums) - 1]
		*chosen = append(*chosen, c)

		helper(nums, chosen, S, count)

		*nums = append(*nums, v)
		*chosen = (*chosen)[:len(*chosen) - 1]
	}
}