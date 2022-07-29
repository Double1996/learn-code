/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
func validPalindrome(s string) bool {
	// 比较简单，
	l, r := 0, len(s) -1
	return check(s,l, r, 0)
}

// 递归的方式
func check(s string, l, r, time int) bool {
		for r > l {  // 两个不相等的时候
		if s[r] != s[l] {
			time++
			if time > 1 {
				return false
			}
			return check(s, l, r-1, 1) || check(s, l+1,r,1)
		}
		l++
		r--
	} 
	return true 
}
// @lc code=end

