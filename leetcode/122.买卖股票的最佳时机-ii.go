/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 贪心的解法
	/*
		收集每天的正利润
	*/
	// var profit int
	// for i:=1; i < len(prices); i++ {
	// 		if prices[i] - prices[i-1] > 0 {
	// 			profit += prices[i] - prices[i-1]
	// 		}
	// }
	// return profit
	
	/**
		使用 动态规划
		定义 dp  dp[i][0] 表示第i天持有的股票现金 dp[i][1] 表示第i天不持有股票的所得最多的现金
	*/

	  dp:=make([][]int,len(prices)) 
    for i:=0;i<len(prices);i++{
        dp[i]=make([]int,2)
    }
    dp[0][0]=-prices[0]
    dp[0][1]=0
    for i:=1;i<len(prices);i++{
        dp[i][0]=max(dp[i-1][0],dp[i-1][1]-prices[i])
        dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i])
    }
    return dp[len(prices)-1][1]	// 最后一天不持有股票的现金
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

