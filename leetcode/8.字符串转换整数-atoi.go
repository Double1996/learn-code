/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	index := 0 
	// 1. 忽略空格
	for index < len(s) && s[index] == ' ' {
			index++
	}


	sign := 1 
	// 2. 判断正负号
	if index<len(s) && (s[index] == '+'||s[index] == '-'){
        if s[index] == '+' {
           sign = 1
           
        }else{
            sign = -1
            
        }
        index++
  }
	ans := 0
	
	for index< len(s) && s[index] >= '0' && s[index] <= '9' {
		if ans > (2147483648 - int(s[index] - '0')) / 10  {
				if sign == 1 {
					return math.MaxInt32
				}
				return math.MinInt32
		}
		ans = ans*10 + int(s[index] - '0')
		index++
	} 
	return ans * sign
}
// @lc code=end

