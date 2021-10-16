/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 动态规划思路
	// 空间
	// 时间
	max := nums[0]
	for i := 1; i  < len(nums); i++ { // 贪心的思路
		if nums[i] + nums[i - 1 ] > nums[i] {
			nums[i] += nums[i -1]
		}
		if nums[i] > max  {
			max = nums[i]
		}
	} 
	return max
}
// @lc code=end

