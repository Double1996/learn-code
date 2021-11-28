package sort

/**
  排序思想；   又称分区交换排序,
  时间复杂读:  最糟糕的情况是 O(n2) 平均是 O(log n)
  空间复杂度:  O(n) 平均是 O(log n)
*/

func partition(arr []int, low, high int) int { // 分区
	index := low - 1
	pivotElement := arr[high] // 预设的最大值
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement { //
			index += 1                              // 第一次 等于 low,
			arr[index], arr[i] = arr[i], arr[index] // 最后的以为，和当前位置的进行交换
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1] // ?
	return index + 1
}

func QuickSortRange(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}
	if low < high {
		pviot := partition(arr, low, high) // 获取
		QuickSortRange(arr, low, pviot-1)  // 排序上半部分
		QuickSortRange(arr, pviot+1, high) //  排序下半部分
	}
}

func QuickSort(arr []int) []int {
	QuickSortRange(arr, 0, len(arr)-1)
	return arr
}
