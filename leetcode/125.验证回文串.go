/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	//  先去除掉特殊的字符
	var sgood string
	 for i := 0; i < len(s); i++ {
        if IsNumber(s[i]) {
            sgood += string(s[i])
        }
    }

    n := len(sgood)
    sgood = strings.ToLower(sgood)  // 或者使用 ascii 运算
		
    for i := 0; i < n/2; i++ {  // 双指针
        if sgood[i] != sgood[n-1-i] {
            return false
        }
    }
    return true
}


func IsNumber(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
// @lc code=end

