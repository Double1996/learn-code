package main

import "fmt"

/**
  排序思想；   又称分区交换排序,
  时间复杂读:  最糟糕的情况是 O(n2) 平均是 O(log n)
  空间复杂度:  O(n) 平均是 O(log n)
*/

func main() {
	var nums = []int{1, 13, 23, 56, 99, 87, 23, 7, 103}

	hasChange := true

	for i := 0; i < len(nums)-1 && hasChange; i++ {
		hasChange = false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				hasChange = true
			}
		}
	}
	fmt.Println(nums)
}
