/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 双指针，假设两链表长度La， Lb，相交部分长度为c
	// 1.先让指针走到链表尾部，改换对方链表头结点, p1走过La, p2走过Lb
	// 2.继续遍历, 直到相交节点，p1走过Lb-c, p2走过La-c
	// 3.p1, p2直到相交节点时，走过的总长度都为La+Lb-c
	// tips: 如果不相交，则两指针各走La+Lb长度，最后都到了各自的尾部，p1=nil
	//for 内循环在两个 链表（a+b和b+a）都结束后就自动结束了。
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 != nil { // 1.走A链表时，走到底 2.改了B链表，继续遍历
			p1 = p1.Next
		} else { // A链表到了尾部，该B头结点
			p1 = headB // 重置 到 headB 
		}
		if p2 != nil { // 与上同理
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1 
    
}
// @lc code=end

