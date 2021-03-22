/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	for i:= len(nums)-1; i > 0; i--  {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i] , nums[i+1:]...)
		}
	}
	return len(nums)
}
// @lc code=end

