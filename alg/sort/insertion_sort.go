package main

import "fmt"

func main() {
	var nums = []int{1, 13, 23, 56, 99, 87, 23, 7, 103}

	for i := 1; i < len(nums)-1; i++ {
		cur := nums[i]
		j := i - 1
		for cur < nums[j] && j > 0 {
			fmt.Println(nums, nums[j+1], nums[j], j)
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = cur
	}
	fmt.Println(nums)
}
