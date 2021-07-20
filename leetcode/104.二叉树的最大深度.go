/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	// 递归写法
	// 时间复杂度： O(n)  树的个数
	// 空间复杂度: o(height) 树的高度
	if root == nil  {
		return 0 
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1 
}

func max(a, b int) int {
	if a> b {
		return a
	} 
	return b 
}
// @lc code=end

