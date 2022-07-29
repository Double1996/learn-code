/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	// 思路
	var res []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {  // 递归写法
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

// 非递归写法, 使用栈
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
  stack := make([]*TreeNode, 0)
  pos := root
  for pos != nil || len(stack) > 0 {
      if pos != nil {
          stack = append(stack, pos)  // 入栈
          pos = pos.Left
      } else {
          node := stack[len(stack)-1] // 出栈
          stack = stack[:len(stack)-1] // 栈递减
          ans = append(ans, node.Val)
          pos = node.Right
      }
    }
    return ans
}

// @lc code=end

