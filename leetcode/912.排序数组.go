/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums) -1) // 传入低位，与 高位的 index
	return nums
}

func quickSort(arr []int, low, high int) {
		if len(arr) <= 1 {
			return
		}
		if low < high {
			pviot := partition(arr, low, high)
			quickSort(arr, low, pviot -1)
			quickSort(arr, pviot+1, high)
		}
}


func partition(arr []int, low, high int) 	int {
		index := low + 1
		pviot := arr[high] 
		for i:=low; i < high; i++ {
			if arr[i] < pviot {
				arr[i], arr[index] = arr[index], arr[i]   // 两个元素进行 swap
				index += 1
			}
		}
		arr[index-1], arr[high] = arr[high], arr[index-1]   // swap
		return index - 1
}


// @lc code=end

