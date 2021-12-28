/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
		// 爬楼梯
		/**
		状态转移方程: dp[n] = dp[n-1] + dp[n-2]
		*/
		if n ==1 {
			return 1
		}

		dp := make([]int, n +1)
		dp[1] = 1
		dp[2] =2
		for i:=3; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
		return dp[n]
}
// @lc code=end

