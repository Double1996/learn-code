/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	// 完成 K笔 交易
	dp := make([][][]int, len(prices))  // 第 i 天， 第 n 是否 拥有股票， 第K次 完成交易

	dp[0][0][0] := 0   // 三维的数组

	for i:= 1; i < len(prices); i++ {
		for j :=1; j < 2;  i++ {   // 
			for c:= 0; c <= k;  c++ {  // 交易次数
					dp[i][0][c] = max(dp[i][1][i-1], dp[k][1][c-1] - prices[i]) 	// 购买股票的状态转移方程
					dp[i][1][c]	 																														// 卖出股票的状态转移方程
			}
		}
	}
	
}


func max(x,y int) int {
	if x > y {
		return x 
	}
	return y
}








// @lc code=end

