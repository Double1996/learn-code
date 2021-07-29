/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	// 满足二叉搜索树， 中序遍历的为顺序遍历则 左 —— 中 - 右边
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max  {
		return false
	}
	return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)  // 接着递归下去
}
// @lc code=end

