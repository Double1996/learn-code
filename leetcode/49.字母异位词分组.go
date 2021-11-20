/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 异位词
	// 思路  使用 hashMap,
	mp := make(map[[26]int][]string)  // 可以使用 字母在 hashmap 表中出现的位置, 作为hash 表的键值
	for _, str  := range strs {   // 遍历字符串
		cnt := [26]int{}
		for _, b := range str {			// 遍历字符串里面的值
			cnt[b - 'a']++   //
		}
		mp[cnt] = append(mp[cnt],  str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans,  v )
	}
	return ans
	// 时间复杂度
	// 空间复杂度
}
// @lc code=end

