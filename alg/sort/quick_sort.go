package main

import "fmt"

/**
  排序思想； 1.
  时间复杂读:  最糟糕的情况是 O(n2) 平均是 O(log n)
  空间复杂度:  O(n) 平均是 O(log n)
*/

func quickSort(array []int) []int {
	return _quickSort(array, 0, len(array)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1) // 左右交换
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

//
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	fmt.Println(quickSort([]int{120, 13, 23, 516, 99, 87, 23, 7, 103}))
}
