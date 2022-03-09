/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
		// 先从从左，再向右
    // 层序遍历，
		res := [][]int{}
		if root == nil {
			return res
		}
		queue := []*TreeNode{root} // 存储节点的队列
		flag := true
		for len(queue) > 0 { // 进行迭代
			  size := len(queue)
				cur := []int{}
				for i:=0; i < size; i++ {
						 node := queue[0] // 获取新的一个节点
						 queue = queue[1:] // 删除第一个元素
						 cur = append(cur, node.Val)
						 if node.Left != nil {
								queue = append(queue, node.Left)
						}
				
						if node.Right != nil {
								queue = append(queue, node.Right)
						}
				}
				if !flag {
						 reverse(cur)
				}
				flag = !flag
				res = append(res, cur)
		}
		return res
}

// 翻转数组
func reverse(a []int) {
	l, r := 0, len(a) -1  // 使用 index 的时候
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
// @lc code=end

