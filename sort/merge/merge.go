package merge

func mergeSort(array []int) {
	if len(array) <= 1 {
		return
	}
	var newArray []int
	mergeHelper(array, newArray, 0, len(array) - 1)
	array = newArray
	return
}

func mergeHelper(array, newArray []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right - left) / 2

	mergeHelper(array, newArray, left, mid)
	mergeHelper(array, newArray, mid + 1, right)

	// set new Array; left copy to right
	//for i := left; i <= right; i++ {
	//	newArray[i] = array[i]
	//}

	// 指针归并
	i, k := left , left
	j := mid + 1
	for i <= mid && j <= right {
		if newArray[i] <= newArray[j] {
			newArray[k] = array[i]
			i++
			k++
		} else {
			newArray[k] = array[j]
			i++
			k++
		}
	}

	// 剩下的有序数据合并
	//if i <= mid {
	//	newArray[k] = array[i]
	//	i++
	//	k++
	//} else {
	//
	//}
}

// 合并 s -> m and m -> e 这两个数组
func merge(array []int, s, m, e int) {
	if array[s] > array[m] {

	}
}
