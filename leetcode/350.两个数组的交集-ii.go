/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	// 时间复杂度: O(mlogm+nlogn)log(n) 
	// 空间复杂度: O(min(m,n))O(min(m,n))
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0
		
	// 交集
	intersect := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intersect = append(intersect, nums1[index1])
			index1++
			index2++
		}
	}
	return intersect
}
// @lc code=end

