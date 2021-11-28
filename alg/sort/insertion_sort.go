package sort

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums)-1; i++ {
		cur := nums[i]
		j := i - 1
		for cur < nums[j] && j > 0 {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = cur
	}
	return nums
}
