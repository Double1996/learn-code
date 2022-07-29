/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
		// 使用单调栈这种数据结构, 维护一个递减的单调栈
		var s []Rect

		if len(height) == 0 {
			return 0
		}
		for _, h := range height {  // 遍历然后确定单调性
			for len(s) != 0 && s[len(s)-1].Height <= h {
					


				// 计算 ans 
			}
		}
}


type Rect  struct {  // 
	Height, Width  int 
}



// 思路二: 双指针的解法












// 思路三：动态规划的解法













// @lc code=end

