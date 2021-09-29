/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	// 迭代的写法
	// 时间复杂度:    
	// 空间复杂度: 
	 var res *ListNode  // 空的链表
	 curr := head				// 

	 for curr != nil {
			next := curr.Next // 1. 把之后所有的节点的存储起来
			curr.Next = res   // 2.  
			res  = curr
			curr = next 
	 }
	 return res

	 // 递归的写法
	//  if head == nil || head.Next == nil {
	// 	 	return head
	//  }
	//  newHead := reverseList(head)
	//  head.Next.Next = head
	//  head.Next = nil 
	//  return newHead
}
// @lc code=end

