/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	stack []int
	minStack []int
}


func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{}}  // 默认一个最小值不是很好理解
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {			// 如果是最小值的情况下，进行
		this.minStack = append(this.minStack, val)
	} else {
			top := this.minStack[len(this.minStack) -1]
			this.minStack = append(this.minStack, min(top, val))
	}
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]  // 删除最后一位
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack) -1]   // 最新的 stack 
}

// 支持最小元素的栈, 额外的信息，肯定需要额外的空间去实现
// 额外的空间去实现
/**
	思路:
	使得每一个元素与其最小的值m, 保持一一对应的。 
	因此我们可以使用一个辅助栈，与元素栈同步插入与删除，用于存储与每个元素对应的最小值
	1.当一个元素要入栈时，我们取当前辅助栈的栈顶存储的最小值，与当前元素比较得出最小值，将这个最小值插入辅助栈中；
  2.当一个元素要出栈时，我们把辅助栈的栈顶元素也一并弹出；
  3.在任意一个时刻，栈内元素的最小值就存储在辅助栈的栈顶元素中。

	方法二： 如何不使用额外的空间去存储
*/
func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack) -1]  // 找出最后一位的 最小栈
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

