/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}  // 迭代需要一个假节点, 主要的作用是为了保证第一个节点可以被遍历到
	for res := dummyHead; res.Next != nil; {
		if res.Next.Val == val {
			res.Next = res.Next.Next  // 链表跳过此节点
		} else {
			res = res.Next    // 正常迭代
		}
	}
	return dummyHead.Next	// 使用迭代的方式
}
// @lc code=end

