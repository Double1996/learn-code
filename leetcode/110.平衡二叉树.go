/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	/**
		思路： 如果是平衡二叉树，则两个 左右子树最大的差值不会大于 1
	*/
	if root == nil {
		return false
	}
	leftDepth, rightDepth := maxDepth(root.Left, root.Right)
	if leftDepth - rightDepth > 1 || rightDepth - leftDepth > 1 {
			return false
	}
	return true 
	// 使用迭代写法
}


// 二叉树的迭代
func maxDepth(root *TreeNode, depthCount int) int {
		if 


}









// @lc code=end

