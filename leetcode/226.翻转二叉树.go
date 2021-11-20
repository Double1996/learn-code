/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
	// 思路
	// 时间复杂度
	// 空间复杂度

	if root == nil {  //de
		return root
	}

		// 交换左右两个节点
		tmp := root.Left
		root.Left = root.Right
		root.Right = tmp
		invertTree(root.Right)
		invertTree(root.Left)
		return root
}
// @lc code=end

