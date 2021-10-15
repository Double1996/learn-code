/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	 plus := 1
    for i := len(digits)-1; i >= 0; i-- {
        digits[i] += plus
        if digits[i]/10 > 0 {
            // 大于等于10时，进位，取模
            plus = digits[i]/10
            digits[i] %= 10
        } else {
            // 小于10时，无需进位，结束
            plus = 0
            break
        }
    }
    if plus > 0 {
        digits = append([]int{1}, digits...)
    }
    return digits
}
// @lc code=end

