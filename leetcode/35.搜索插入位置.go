/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	// 目标值的问题
	// 二分查找 要使用 递归的方式, 但是有点变形
	// 这道题 非常的巧妙
	
	left, right := 0, len(nums) -1
	ans := len(nums) // 定义一个最右边的位置

	for left <= right {
		mid := left + (right - left) >> 1 
		if nums[mid] < target {
			 left = mid +1
		} else if nums[mid] >= target {
			 ans = mid 
			 right = mid -1 	
		}
	}
	return ans
}	
// @lc code=end

