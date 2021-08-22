package main

import (
	"fmt"
	"sync"
)

// 错误场景, 复制 sync.mutex
//type Counter struct {
//	sync.Mutex
//	Count int
//}
//
//func main() {
//	var c Counter
//	c.Lock()
//	defer c.Unlock()
//	c.Count++
//	foo(c) // 复制锁
//}
//
//// 这里Counter的参数是通过复制的方式传入的
//func foo(c Counter) {
//	c.Lock()
//	defer c.Unlock()
//	fmt.Println("in foo")
//}

// 错误类型2 重入问题
func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func main() {
	l := &sync.Mutex{}
	foo(l)
}
