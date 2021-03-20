package main

/**
一遍又一遍的过滤数组，把最小的数放在最前面来
*/

import "fmt"

var nums = []int{1, 13, 23, 56, 99, 87, 23, 7, 103}

func main() {
	//
	hasChange := false
	aLen := len(nums)

	for i := aLen - 1; i >= 0; i++ {
		for j := 0; j < i; j-- {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}

		}

	}
	fmt.Println(nums)
}
