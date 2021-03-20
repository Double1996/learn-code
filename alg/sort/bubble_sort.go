package main

import "fmt"

func main() {
	var nums = []int{1, 13, 23, 56, 99, 87, 23, 7, 103}

	hasChange := true

	for i := 0; i < len(nums)-1 && hasChange; i++ {
		hasChange = false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1] = nums[j]
				hasChange = true
			}
		}
	}
	fmt.Println(nums)
}
