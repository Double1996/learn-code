package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 时间复杂度
	// 空间复杂度
	// 步骤

	sort.Ints(nums) // 排序
	var res [][]int

	l := len(nums)

	for i := 0; i < l-2; i++ { // 双指针的选择方式
		n1 := nums[i]
		if n1 >= 0 { // 边界条件
			break
		}
		if i > 0 && nums[i-1] == nums[i] { // 去重
			continue
		}

		a := i + 1
		b := l - 1

		for a < b {
			n2, n3 := nums[a], nums[b]
			if n1+n2+n3 > 0 {
				b--
			} else if n1+n2+n3 < 0 {
				a++
			} else {
				res = append(res, []int{i, a, b})
				for a < b && nums[a] == n2 {
					a++
				}
				for a < b && nums[b] == n3 {
					b--
				}
			}
		}
	}
	return res
}

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}
