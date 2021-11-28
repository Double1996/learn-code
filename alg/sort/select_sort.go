package sort

/**
思路：一遍又一遍的过滤数组，把最小的数放在最前面来
*/

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := 0; j < len(arr); j++ {
			if arr[j] < arr[i] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
