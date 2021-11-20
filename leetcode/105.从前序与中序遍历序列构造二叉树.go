/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 思想: 一段 preOrder 一段 inorder 去进行还原的问题
	// 递归的思想
	
	root := &TreeNode{
		Val: preorder[0],
	}
	return build(0, len(preOrder) -1, 0, inorder-1)
}

// 用preOrder[l1, r1] 与 inOrder[l2, r2] 来还原二叉树 
func build(l1, r1, l2, r2 int) *TreeNode{
		


}
// @lc code=end

