/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	// 逆波兰表达式 就是后缀表达式
	// 遇到操作符，就拿最近的结果 去做运算
	// 时间复杂度 O(n)
	// 空间复杂度 O(n)

	stack := []int{}  // 需要额外的空间 去存储运算的结果

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token  == "/" {
			// 抓取相邻的两个值，做 一个 结果
			y := stack[len(stack) -1]
			stack = stack[:len(stack) -1]

			x := stack[len(stack) -1]
			stack = stack[:len(stack) -1]

			stack = append(stack, cal(x, y, token))
		} else {
			// 不为特殊符号是直接入栈的操作
			tokenInt, _ := strconv.Atoi(token)
			stack = append(stack, tokenInt)	
		}
	}
	return stack[0]
}

// 对于值的一个计算
func cal(x, y int, token string) int {
		switch token{
			case "*":
				return x * y
			case "/":
				return x / y
			case "+":
				return x + y
			case "-":
				return x - y
			default:
				return 0 		
		}	
}

// @lc code=end

