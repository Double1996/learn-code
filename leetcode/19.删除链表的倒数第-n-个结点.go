/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	/**
	 思路:  设置哑节点，前后双指针的方式, 快慢指针的方式
	 时间复杂度: O(L)
	 空间复杂度； O(1)
	*/
	dummy := &ListNode{0, head} // 新建一个哑节点
	first, second := head, dummy
	for i := 0; i < n; i ++ {
		first = first.Next    // 先跑 快指针
	}
	
	 for ;first != nil; first = first.Next {
        second = second.Next
    }
		second.Next = second.Next.Next  // 删除中间的节点

		return dummy.Next
}
// @lc code=end

