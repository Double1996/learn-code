/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// 1. 定义 dp 结果
	dp := make([]int, len(nums))
	// 2. f(n) 状态转移的方程 
	// 3. 初始化, 每一个 dp[i] 为 1
	// 4. 怎样遍历都是可以的
	var ans int
	for i, _ := range nums {
		dp[i] = 1
	}

	for i, _ := range nums {
		for j:=0; j < i; j++ {
			if i > 0 && nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] +1) // 状态转移方程
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return  ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

