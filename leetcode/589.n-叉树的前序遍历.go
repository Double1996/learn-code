/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// var res []int

func preorder(root *Node) []int {
	 // 思路： 使用栈来进行 深度遍历
	 /**
	 我们首先把根节点入栈，因为根节点是前序遍历中的第一个节点。
	 随后每次我们从栈顶取出一个节点 u，它是我们当前遍历到的节点，
	 并把 u 的所有子节点逆序推入栈中。
	 例如 u 的子节点从左到右为 v1, v2, v3，那么推入栈的顺序应当为 v3, v2, v1，
	 这样就保证了下一个遍历到的节点（即 u 的第一个子节点 v1）出现在栈顶的位置。
	 */
	 res = []int{}
	 stack := []*Node{root}  // 使用栈来实现递归
	 if len(stack) != 0 {  // 栈不为空的情况下
		for root != nil {
			res := append(res, res.Val)
			if len(root.Children) == 0 {  // 如果没有子元素
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- {   // 按照 先右边后左边的顺序放入
				stack = append(stack, root.Children[i]) //入栈
			}
			root = root.Children[0]
		}
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	 }	
	//  dfs(root)
	 return res 
}

// func dfs(root *Node) {   // 深度优先遍历, 递归实现
// 	if root == nil {
// 		return
// 	}
// 	res = append(res, root.Val)
// 	for _, n := range root.Children{
// 		dfs(n)
// 	}
// }

// @lc code=end

