/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	// 二分查找的
	left, right := 0, x
	var ans int	
	for left <= right { // 注意这里的
		mid := left + (right - left) / 2
		if mid * mid > x {
				right = mid -1
		} else if mid * mid == x {
				return mid
		} else {
				ans = mid
				left = mid + 1
		}
	}
	return ans
}
// @lc code=end

