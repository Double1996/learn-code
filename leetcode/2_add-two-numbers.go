package leetcode

/**
  看题目的描述一脸懵逼
  逆序:
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		cur.Next = new(ListNode) // 遍历操作
		cur = cur.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Val = carry % 10
		carry /= 10
	}
	return dummy.Next
}
