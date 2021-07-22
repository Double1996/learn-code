/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := list.New()   // 定义一个队列
	queue.PushBack(root)
	var tempArr []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i:= 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)	
			if node.Left != nil {
				queue.PushBack(node.Left)
			}	
			if node.Right != nil {
				queue.PushBack(node.Right)
			}	
			tempArr = append(tempArr, node.Val)
		}		
		res = append(res, tempArr)		
		tempArr = []int{}
	}
	return res
}
// @lc code=end

