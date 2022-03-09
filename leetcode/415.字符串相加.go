/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	// 双指针，逆序-进位-反转
	/*
	我们定义两个指针 ii 和 jj 分别指向 \textit{num}_1num 
  和 \textit{num}_2num 
  的末尾，即最低位，同时定义一个变量 \textit{add}add 维护当前是否有进位，然后从末尾到开头逐位相加即可。
	你可能会想两个数字位数不同怎么处理，这里我们统一在指针当前下标处于负数的时候返回 00，等价于对位数较短的数字进行了补零操作
	*/
	var res string
	a, b, ca := len(num1) -1, len(num2) -1, 0 // 左右的指针，进位的
	var n1, n2, sum int
	for a >=0 || b >= 0 {  // 从大到小的进行遍历，
		n1, n2 = 0, 0
		if a >= 0 {
			 n1 = int(num1[a] - '0')
		}
		if b >=0 {
			  n2 = int(num2[b] - '0')
		}
		// 该位权数和
		sum =  n1 + n2 + ca
		// 叠加进位
		ca = sum / 10  // 最高进位

		// 每次的res都是逆序的，所以最后是正序
		res = strconv.Itoa(sum%10) + res
		a--          // 递减的操作
		b--
	}
	if ca > 0 {
		return "1" + res
	}
	return res

}
// @lc code=end

