package main

import "fmt"

func main() {
	list := []int{8, 2, 4, 1, 5, 3, 7, 3}
	sort(list, 0, len(list) - 1)
	fmt.Printf("%+v", list)
}

// 从 i、j 开始查找交换坑，直到 i == j 时
// 左边都是比 x 小的右边都是比 x 大的
func sort(a []int, n, m int) {
	// 退出条件
	if n > m {
		return
	}
	// 取第一个基准数
	t, i, j := a[n], n, m
	for i < j {
		// 从右向左找第一个小于 t 的数，放入 a[i]
		// 因为左边的数都需要比 t 小，我们取出的是第一个数
		for i < j  {
			// 找到一个数
			if a[j] <= t {
				a[i] = a[j]
				break
			}
			// 越过的数都比 t 大
			j--
		}

		// 从左向右找第一个大于 t 的数，放入 a[j]
		for i < j {
			if a[i] > t {
				a[j] = a[i]
				break
			}
			i++
		}
	}
	// i == j 写回基准数
	a[i] = t
	sort(a, n, i - 1)
	sort(a, i + 1, m)
	return
}