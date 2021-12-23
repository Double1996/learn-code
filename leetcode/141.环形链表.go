/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 时间复杂度  O(N)		// 当链表中不存在当链表中不存在环时，快指针将先于慢指针到达链表尾部，链表中每个节点至多被访问两次。
	// 空间复杂度  O(1)   // 只使用了两个指针的额外空间
	// 思路: 快慢指针, 判断一个链表是否有环, 有环发生相遇，无环不会相遇, fast 走 两步， slow 
	// 走 一步
	
	if head == nil {
		return false
	}

	fast,slow := head, head 
	for fast.Next != nil && fast.Next.Next != nil {  // for 循环的判断不能写反了
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				return true
			}
	}
	return false
}

// @lc code=end

