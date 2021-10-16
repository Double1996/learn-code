/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] 统计「优美子数组」
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	// 子段的数量的 n2
	// 原文 奇数 是1,  偶数 为 1 (每个数 mod2), 
	// 简化成 两数之差问题, 子段和是K的数据 s[r] - s[j -1] = k 
	// 前缀和
	var arr = make([]int, len(nums) + 1)
	arr[0] = 1  // 
	res, sum := 0, 0
	for _, num := range nums {
		sum += num % 2	 // 
		arr[sum]++
		if sum >= k {
			res += arr[sum -k]
		}
	}
	return res
}
// @lc code=end

