/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	// 思想 分治
	// 时间复杂度:
	// 空间复杂度:
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1/x
	}
	temp := myPow(x, n / 2)
	ans := temp * temp
	if ( n % 2) == 1 {

	}
	return ans
}
// @lc code=end

