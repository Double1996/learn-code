/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {

// 中缀表达式， 与 基本计算器1 相比 没有括号 比较简单
// 比较难的点： 遇到符号的时候，是否 需要进行运算。 查看最近相关性
// 时间复杂度，空间复杂度都是 O(n)

   stack := []int{}
   preSign := '+'
	 num, ans := 0, 0 
    for i, ch := range s {
        isDigit := '0' <= ch && ch <= '9'
        if isDigit {
            num = num*10 + int(ch-'0')
        }
        if !isDigit && ch != ' ' || i == len(s)-1 {
            switch preSign {
            case '+':
                stack = append(stack, num)
            case '-':
                stack = append(stack, -num)
            case '*':
                stack[len(stack)-1] *= num
            default:
                stack[len(stack)-1] /= num
            }
            preSign = ch
            num = 0
        }
    }
    for _, v := range stack {
        ans += v
    }
    return ans
}
// @lc code=end
