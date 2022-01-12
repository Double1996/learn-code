/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	// 使用 额外的线性链表来实现
	if  head == nil {
		return 
	}	
	nodes := []*ListNode{}
	
	for node := head; node != nil; node = node.Next {
			nodes = append(nodes, node)
	}
	
	// 使用双指针
	i, j := 0, len(nodes) -1
	for i < j {
		nodes[i].Next = nodes[j]
		i ++ 
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil

	
}
// @lc code=end

