/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false // 判断是否是成对出现的
	}
  pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
	stack := []byte{}
	for i := 0; i < n; i++ {
			if pairs[s[i]] > 0 {
					if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
							return false
					}
					stack = stack[:len(stack)-1]
			} else {
					stack = append(stack, s[i])
			}
	}
  return len(stack) == 0
}
// @lc code=end

