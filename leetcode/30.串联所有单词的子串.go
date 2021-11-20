/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	// 时间复杂度
	// 空间复杂度
	// 使用的知识点: hash 表, 怎么判断 两个
	

  l := len(words) 
	var res []int
	if l == 0 {  // 边界条件
		return res 
	}

	var tot int
	m := make(map[string]int, len(words))    // 统计每一个单次出现的次数
	for _,  word := range words {
			tot += len(word)
			if v, ok := m[word]; ok {
				v++ 
			}
	}
  for  i=0; i + tot < len(s); i ++ {
		 if (vaild(string(s[i:tot]), words, m) {
			res = append(res, i)			 
		 }
	}
	return res
}

func vaild(subStr string,  words []string, m map[string]int) bool {  // 怎么子字符串, 是否
	subStrm :=  make(map[string]int, len(words))    // 统计每一个单次出现的次数
	workL := len(words[0])
	for i:=0; i + wordL < len(subStr); i + wordL {
		subStrm[string(subStr[i:i+wordL])]++  
	}
	return equlesMap(subStrm, m)  // 比较两个map 是否相等
}

func equlesMap(a,b map[string]int)  bool { 	// 比较两个 map 是否相等
	for k,av := a {
			if bv,ok := b[k]; !ok {
				return ok
			}
	}
	return true
}
// @lc code=end

