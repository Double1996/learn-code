/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
		// 层序遍历的时候层序遍历的时候，判断是否遍历到单层的最后面的元素，
		// 如果是，就放进result数组中，随后返回result就可以了。
	 //  还是的 层序遍历的模板
  
	 res := [][]int{} // 每一层遍历的结果
	 finalRes := []int{}
	 if root == nil {
		 return finalRes
	 }
	 
	 queue := []*TreeNode{root}
	 for len(queue) > 0 {
			size := len(queue)
			curr := []int{}
			for i := 0; i < size; i++ {
					node := queue[0]
					queue = queue[1:]  // pop 操作
					curr = append(curr, node.Val)
					if node.Left != nil {
					 	queue = append(queue, node.Left)				
					}

					if node.Right != nil {
					 queue = append(queue, node.Right)				
					}
			}
			res = append(res, curr) // 加入当前的的最后一个
	 }

	 for _, arr := range res {
			finalRes = append(finalRes, arr[len(arr) -1]) // 最后一个元素
	 }
	 return finalRes
}
// @lc code=end

