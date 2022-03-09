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

    arr := []*TreeNode{root}

    for len(arr) > 0 {
        size := len(arr)
        curRes := []int{}
        for i := 0; i < size; i++ {
            node := arr[i]
            curRes = append(curRes, node.Val)
            if node.Left != nil {
                arr = append(arr, node.Left)
            }
            if node.Right != nil {
                arr = append(arr, node.Right)
            }
        }
        arr = arr[size:]
        res = append(res, curRes)
    }

    return res
}
// @lc code=end

