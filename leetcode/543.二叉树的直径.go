/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
var ans int
func diameterOfBinaryTree(root *TreeNode) int {
	 // 思路: 最大的直径 = 左子树的高度 + 右边子树的高度
	 ans = 0
	 depthRoot(root)
	 return ans		
}


func depthRoot(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftNodeDepth := depthRoot(node.Left)
	rightNodeDepth := depthRoot(node.Right)
	if leftNodeDepth+rightNodeDepth > ans {
		ans  = leftNodeDepth + rightNodeDepth  // 如果发现最大的值就进行更新
	}
	return max(leftNodeDepth, rightNodeDepth) + 1  // 不行的就返回当前的深度
}


func max(x, y int) int {
	if x >y  {
		return x
	}
	return y

}

// @lc code=end

