/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	/**
	由于给定的链表是排好序的，因此重复的元素在链表中出现的位置是连续的，
	因此我们只需要对链表进行一次遍历，
	就可以删除重复的元素。由于链表的头节点可能会被删除，
	因此我们需要额外使用一个哑节点（dummy node）指向链表的头节点。

  具体地，我们从指针 \textit{cur}cur 指向链表的哑节点，随后开始对链表进行遍历。
	如果当前 \textit{cur.next}cur.next 与 \textit{cur.next.next}cur.next.next 对应的元素相同，
	那么我们就需要将 \textit{cur.next}cur.next 以及所有后面拥有相同元素值的链表节点全部删除。
	我们记下这个元素值 xx，随后不断将 \textit{cur.next}cur.next 从链表中移除，
	直到 \textit{cur.next}cur.next 为空节点或者其元素值不等于 xx 为止。
	此时，我们将链表中所有元素值为 xx 的节点全部删除。

  如果当前 \textit{cur.next}cur.next 与 \textit{cur.next.next}cur.next.next 对应的元素不相同，
  那么说明链表中只有一个元素值为 \textit{cur.next}cur.next 的节点，
  那么我们就可以将 \textit{cur}cur 指向 \textit{cur.next}cur.next。

	*/

	if head == nil {
		return head
	}
	// 使用哑节点
	dummpy := &ListNode{Next: head}
	curr := dummpy // 重新复制一份数据
	for curr.Next != nil && curr.Next.Next != nil {
			tmp := 0
			if curr.Next.Val == curr.Next.Next.Val {
					tmp = curr.Next.Val
					for curr.Next != nil && curr.Next.Val == tmp {
							curr.Next	= curr.Next.Next  // 强制下一步迭代
					}
			} else {
					curr	= curr.Next
			}
	}
	return dummpy.Next
}
// @lc code=end

