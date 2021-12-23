/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
//     /**
//         1. 先分组
//         2. 内部要翻转, 调用翻转链表
//         3. 前面和后面的边需要
//     **/
//     hair := &ListNode{Next: head}   // 定义一个结果, head 之前是 hair 
//     pre := hair                     // 这个是 前一个链表

//     for head != nil {               // 进行链表的遍历
//         tail := pre                 // 新增一个 尾巴的 临时的 链表
//         for i := 0; i < k; i++ {    // 先进行分组
//             tail = tail.Next        // i < K 进行跳过 分组, tail 尾巴 要跳过进行遍历
//             if tail == nil {        // 如果 尾巴 直接为空,  那么直接返回结果
//                 return hair.Next    // 直接返回结果
//             }
//         } 
//         nex := tail.Next                    // 保存 目前 翻转前 尾巴部分的指针
//         head, tail = myReverse(head, tail)  // 翻转

//         pre.Next = head
//         tail.Next = nex

//         pre = tail
//         head = tail.Next
//     }
//     return hair.Next   // 返回空节点的 Next
// }

// // 一段节点内的翻转
// func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
//     prev := tail.Next   // ?
//     p := head          // 拷贝一份 链表的头  
//     for prev != tail {   // 
//         nex := p.Next    // nex  P 的 Next 指针
//         p.Next = prev       
//         prev = p
//         p = nex         
//     }
//     return tail, head

    /**
        第二种写法 极客时间的写法
        分组遍历的 框架
    */
    dummy := &ListNode{0, head}  // 必须分成两部分走
    last := dummy
    for head != nil {
        end := getEnd(head, k)
        if end == nil {
            break
        }
        next := end.Next  // 1. 保存 next, 分组
        
        // 2. 一组内部去翻转整个链表
        reverse(head, next)

        // 3. 更新前一组, 后一组的边
        last.Next = end
        head.Next = next

        last = head
        head = next      // 分组的head
    }   
    return dummy.Next
 }

// 用于翻转 head 与 end 之间的 链表
func reverse(head, end *ListNode) {
    var res *ListNode = head   // 
    for head != end  {   // 翻转链表相同的代码
        next := head.Next
        head.Next = res
        res = head       
        head = next
    } 
}

func getEnd(head *ListNode, k int) *ListNode {
    for head != nil {
        k--
        if k == 0 {
            return head
        }    
        head = head.Next       
    }    
    return nil
}

// @lc code=end

