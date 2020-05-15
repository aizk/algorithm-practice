package main

import "fmt"

// 合并两个有序数组
func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}

	fmt.Printf("%+v", merge(x, y, len(x), len(y)))
	return
}

func merge(x, y []int, lx, ly int) []int {
	i := lx - 1
	j := ly - 1
	// x 写入位置
	index := lx + ly - 1
	x = append(x, y...)

	for i >= 0 && j >= 0 {
		if x[i] >= y[j] {
			x[index] = x[i]
			i--
		} else {
			x[index] = y[j]
			j--
		}
		index--
	}
	if i > j {
		return x
	} else {
		for j >= 0 {
			x[index] = y[j]
			index--
			j--
		}
	}
	return x
}