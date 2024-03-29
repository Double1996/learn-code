/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    // prehead := &ListNode{}
    // result := prehead       // 哨兵节点
    // for l1 != nil && l2 != nil {
    //     if l1.Val < l2.Val {
    //         prehead.Next = l1
    //         l1 = l1.Next
    //     }else{
    //         prehead.Next = l2
    //         l2 = l2.Next
    //     }
    //     prehead = prehead.Next
    // }
    // if l1 != nil {
    //     prehead.Next = l1
    // }
    // if l2 != nil {
    //     prehead.Next = l2
    // }
    // return result.Next

     dummpy := &ListNode{}
     res := dummpy  // 哨兵节点
     for list1 != nil && list2 != nil {
         if list1.Val > list2.Val {
            dummpy.Next = list2
            list2 = list2.Next
         } else {
            dummpy.Next = list1
            list1 = list1.Next
            
         }
         dummpy = dummpy.Next // 虚拟节点进行迭代
     }
     if list1 == nil {
         dummpy.Next = list2
     } 
     if list2 == nil {
         dummpy.Next = list1
     }
     return res.Next   
}
// @lc code=end

