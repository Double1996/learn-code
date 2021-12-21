/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	NodeQ := []*Node{root} // 使用队列来进行广度优先的遍历
	res := [][]int{}				// 

	for len(NodeQ) == 0 {  
		if root != nil {
			for i:= 0; i < len(root.Children); i++ {
				


			}
		}	
	}
}
// @lc code=end

