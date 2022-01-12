/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
		// 思路 最大 整数 ans * ans = x
		// 使用 二分法 来求解
		l, r := 1, x
		var ans int
		for l <= r {
			  mid := l + (r - l ) /2  // 二分模板寻找中间的 mid
				if mid * mid <= x  {
						ans = mid
						l = mid + 1					
				} else {
						r = mid -1 
				}
		}
		return  ans
}
// @lc code=end

