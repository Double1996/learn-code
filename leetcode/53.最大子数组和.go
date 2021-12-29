/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 使用动态规划的思想来解决
	dp := make([]int, len(nums))  // 定义 dp 的含义
	ans := nums[0]
	// 初始化 每个 数组都是 自己的本身
	for index, num := range nums {
		dp[index] = num
	}

	// 
	for i:=1; i<len(nums); i++ {
			dp[i] = max(dp[i-1] + nums[i], nums[i])
			if dp[i] > ans {
				ans = dp[i]
			}
	}
	return ans
}


func max (x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

