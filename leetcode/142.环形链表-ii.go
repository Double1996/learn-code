/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 但是 双指针
 fast, slow := head, head // 先定义一个双指针
 hasCycle := false   			// 判断是否有环

 for fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
			  hasCycle = true	
				break 
			}
 }

 if !hasCycle {   // 没有环的情况下就等于空
	 return nil 
 }

slow = head
for slow != fast {   // 开始从头
	slow = slow.Next
	fast = fast.Next
}


	return fast
}

// 第二种思路，就是使用 hashMap https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/4ms-37mb-by-1990jxy-rvvd/


// @lc code=end

