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

	 for curr != nil {    //  递归的思想, 如果没有为空改变 链表的指向
			next := curr.Next // 1. 把之后所有的节点的存储起来
			curr.Next = res   // 2. 当前的链表 指向为结果链表
			res  = curr       // 3. 结果等于 当前链表
			curr = next				// 4. 链表 迭代完成 
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

