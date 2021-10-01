/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
		// 双指针的思路
		// 时间复杂度  Olog(n)
		// 空间复杂度  2n

		var a,b int
		sorted := make([]int, 0, m+n)
		
		for  {
			if a == m {
				sorted = append(sorted, nums2[b:]...)
				break
			}

			if b == n {
				sorted = append(sorted, nums1[a:]...)
				break
			}
			
			if nums1[a] < nums2[b] {
				sorted= append(sorted, nums1[a])
				a++
			} else {
				sorted= append(sorted, nums2[b])
				b++			
			}
		}
		copy(nums1, sorted)
}
// @lc code=end

