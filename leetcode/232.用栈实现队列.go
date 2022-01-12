/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	inStack, outStack []int
}


func Constructor() MyQueue {
	return MyQueue{}
}

// 添加元素
func (this *MyQueue) Push(x int)  {  
	this.inStack = append(this.inStack, x)  // 添加 末尾
}

func (this *MyQueue) in2Out() {
	for len(this.inStack) > 0 {  // 这样慢慢的抽取，就可把队首 放到 outstack 最后一位上面
		this.outStack = append(this.outStack, this.inStack[len(this.inStack) -1])
		this.inStack = this.inStack[:len(this.inStack) -1]
	}	
}

// 获取开头的一个元素
func (this *MyQueue) Pop() int {   // 
	if len(this.outStack) == 0 {
		this.in2Out()
	}
	x := this.outStack[len(this.outStack) -1 ]  // 最后一位
	this.outStack = this.outStack[:len(this.outStack)-1]
	return x
}

func (this *MyQueue) Peek() int {   // 取队首
	if len(this.outStack) == 0 {
		this.in2Out()
	}
	return this.outStack[len(this.outStack) -1 ]  // 取第一位
}


func (this *MyQueue) Empty() bool {  // 判断是否为空
	return len(this.inStack) == 0 && len(this.outStack) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

