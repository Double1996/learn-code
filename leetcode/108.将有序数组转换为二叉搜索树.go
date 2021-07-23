/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
		if len(nums) == 0 {
		return nil
	}

	n := (len(nums) - 1) / 2

	var root = &TreeNode{}
	root.Val = nums[n]

	root.Left = sortedArrayToBST(nums[:n])  // 

	root.Right = sortedArrayToBST(nums[n+1:])

	return root
}
// @lc code=end

