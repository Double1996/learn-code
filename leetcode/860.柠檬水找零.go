/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	// 简单的逻辑使用贪心。
	// 连个用来计数
	// 时间复杂度 O(n) 空间 O(1)  // 空间常量
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 3 {
				five -=3
			} else {
				return false
			}
		}
	}
	return true 
}

// @lc code=end

