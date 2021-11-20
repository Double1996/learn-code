/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

var res []string
func letterCombinations(digits string) []string {
	/**
		思路: 进行搜索, 每一个选项里面有3到4个选项，也就是从 n3 到 n4 中选择进行指数型的搜索
		使用深度优先搜索

	
		时间复杂度:
		空间复杂度:
	*/
	 res := []string
	 // 定义字符数字





	 return res
}

func dfs(index int , str string) {	 // 两个参数变化的
	if index == len(digits) {
			// 将答案存储进来
			res = append(res, str)  
			return
	}
	for   



}


// @lc code=end

