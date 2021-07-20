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
	 res := &ListNode{}
	 curr := head

	 for curr != nil {
			next := curr.Next
			curr.Next = res
			res  = curr
			curr = next 
	 }
	 return res

	 // 递归的写法









}
// @lc code=end

