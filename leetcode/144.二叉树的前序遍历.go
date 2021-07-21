/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
		// 前序遍历是 根左右
		// 思路
		// 时间复杂度:  O(n) 
		// 空间复杂度:  O(n)
		var vals []int
		var preorder func(*TreeNode)  //  必须先定义 一个空壳的func
		preorder = func(node *TreeNode) { // 定义一个子函数
			if node == nil {  // 
				return
			}
			vals = append(vals, node.Val)
			preorder(node.Left)
			preorder(node.Right)
		}
		preorder(root) 				
		return vals 
}
// @lc code=end

