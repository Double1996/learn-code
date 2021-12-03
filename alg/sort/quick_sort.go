package sort

/**
  排序思想；   又称分区交换排序,
  时间复杂读:  最糟糕的情况是 O(n2) 平均是 O(log n)
  空间复杂度:  O(n) 平均是 O(log n)
  写法递归： 主要是两个函数 1. 排序一个数组的 低点 与 高点, 查询 数组 数据的
                        2.
  两个函数参数都是一样的， 传递数组，
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
	arr[index+1], arr[high] = arr[high], arr[index+1] // swap low and high
	return index + 1
}

func QuickSortRange(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}
	if low < high {
		pviot := partition(arr, low, high) // 获取基点的位置, 并且排序好，大于基点的放在右边，小于的放在左边
		QuickSortRange(arr, low, pviot-1)  // 排序上半部分
		QuickSortRange(arr, pviot+1, high) // 排序下半部分
	}
}

func QuickSort(arr []int) []int {
	QuickSortRange(arr, 0, len(arr)-1)
	return arr
}
