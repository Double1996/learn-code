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
  pairs := map[byte]byte{ // 定义 括号的hash, key 都是关闭的
        ')': '(',
        ']': '[',
        '}': '{',
    }
	stack := []byte{}
	for i := 0; i < n; i++ {
			if pairs[s[i]] > 0 { // 如果命中
					if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
							return false // 如果当前的栈是空，或者 栈的最后和对应的 pair[s[i]] 不符合
					}
					stack = stack[:len(stack)-1] // 出栈
			} else {
					stack = append(stack, s[i])
			}
	}
  return len(stack) == 0 // 所用的栈的都出完了, 判断栈是否清空掉
}
// @lc code=end

