/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if v, ok := m[target-n]; ok {
			return []int{i, v}
		}
		m[n] = i
	}
	return nil
}
// @lc code=end

