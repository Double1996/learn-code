/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	/**
		动态规划五部曲: 
		1. 确定 dp 数组和下标的含义: 筹好总额为j 所需钱币的最少个数 dp[j]
		2. 确定 dp 递推公式 dp[j] = min(dp[j - coins[i]] + 1, dp[j]);
		3. dp 数组如何 进行初始化
		4. 确定 遍历 顺序
	*/
	// 先遍历物品再遍历背包
	dp := make([]int, amount +1) // ?
	dp[0] = 0
	
	// 初始化 j
	for j:=1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 先 遍历物品
	for i := 0; i < len(coins); i++ { 
		// 再遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j - coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j -coins[i]] + 1)
			}
		}
	}
	
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x 
}
// @lc code=end

