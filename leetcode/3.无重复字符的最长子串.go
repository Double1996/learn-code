/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
// func lengthOfLongestSubstring(s string) int {
//     m := map[byte]int{} // 记录到每一个字符串的出现的 int 
//     n := len(s)
//     // 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
//     rk, ans := -1, 0
//     for i := 0; i < n; i++ {
//         if i != 0 {
//             // 左指针向右移动一格，移除一个字符
//             delete(m, s[i-1])
//         }
//         for rk + 1 < n && m[s[rk+1]] == 0 {
//             // 不断地移动左右指针
//             m[s[rk+1]]++
//             rk++
//         }
//         // 第 i 到 rk 个字符是一个极长的无重复字符子串
//         ans = max(ans, rk - i + 1)
//     }
//     return ans
// }

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func lengthOfLongestSubstring(s string) int {
    // 滑动窗口的思路
    ans := 0
    left, right := 0, 0
    for right < len(s) {
       for i:= left; i < right;  i++ {
           if s[i] == s[right] {
               ans = max(ans, right - left)
               left = i + 1
               break
           }
       }
       right++
    }
    ans = max(ans, right - left)
    return ans
}
// @lc code=end

