/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	// 时间复杂度：
	// 空间复杂度:
	// 解法:  双指针 二分查找
	var res []int
	l := 1    // 双指针解法
	r := len(numbers)
	for l < r {
		count := numbers[l-1] + numbers[r-1]
		if count == target {
			res = []int{l, r}
			return res
		} else if count > target {
			 r--
		} else if count < target {
			l++
		} 
	}
	return res
}
// @lc code=end

