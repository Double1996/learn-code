package 第一周

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 时间复杂度 O(n+m)  遍历了两个链表
	// 空间复杂度 O(1)

	prehead := &ListNode{} // 哨兵节点
	res := prehead
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			prehead.Next = l2
			l2 = l2.Next
		} else {
			prehead.Next = l1
			l1 = l1.Next
		}
		prehead = prehead.Next // 将 哨兵节点向后位移
	}
	if l1 == nil { // 如果其中的一个链表为空, 链表指向另外一个链表
		prehead.Next = l2
	}

	if l2 == nil {
		prehead.Next = l1
	}

	return res.Next
}
