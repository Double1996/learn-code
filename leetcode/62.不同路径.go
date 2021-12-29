/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// 不同路径
	// dp 五部曲：
	/**
	1. dp[i][j] ：表示从（0 ，0）出发，到(i, j) 有dp[i][j]条不同的路径。
	2. 递推公式：dp[i][j] 表示 dp[i-1][j] , dp[i][j-1]
	3. 如何初始化 dp 数据,  到边缘的数据都是 1
	4. 如何 确定 dp 的遍历顺序

	*/
	dp := make([][]int, m)  // 第一步, 确定数组的数据的含义

	for i := range dp {   // 初始化 dp的数据
		dp[i] = make([]int, n)  // 初始化所有的数组
		dp[i][0] = 1				    // 初始化数据 为1
	}

	for j:= 0; j < n; j++ { // 初始化 dp 的数据
		dp[0][j] =1           // 初始化 
	}

	for i:=1; i < m; i++ {   
		for j:=1; j< n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] // f(n) //
		}
	}
	return dp[m-1][n-1]   // 这个最后的数据 不知为什么
}
// @lc code=end

