/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 简单的贪心实现
	// 暴力的算法不可取，去最小的
	var maxProfit int
	minPrice := math.MaxInt32
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if  price - minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

// @lc code=end

