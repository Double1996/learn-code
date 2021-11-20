/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	left, right  := 0, len(nums) -1  // 寻找最小的值
	for left < right {
			mid := (left +  right) / 2  // 寻找到中间的值
			if nums[mid] <= nums[right] {   // 
				right = mid
			} else {
				left = mid + 1 				
			}
	}
	return nums[right]
	// 思路  二分法
	// 时间复杂度：
	// 空间复杂度：
}
// @lc code=end

