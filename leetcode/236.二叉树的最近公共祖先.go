/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//  使用递归三部曲
  //  使用到回溯的算法思想: 先进性

	if root == nil {
		return root 
	}

	if root == p ||  root == q {
		return root
	}
		
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)		

	// 左右都不为空
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	
	if right != nil {
		return right
	}
 
	return nil 
}
// @lc code=end

