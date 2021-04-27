package main

import (
	"fmt"
	"time"
)

/**
panic 只会触发goroutine 里面的延迟函数调用
*/
func func1() {
	defer println("in main")
	go func() {
		defer println(" in goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)
}

// defer 是 后进先出
func func2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// go 里面是进行值传递的, 得到的结果为0
func func3() {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt))

	time.Sleep(time.Second)
}

// 可以通过匿名函数来进行, 来获取函数执行的时间间隔
func func4() {
	startedAt := time.Now()
	defer func() { // 拷贝的函数的指针
		fmt.Println(time.Since(startedAt))
	}()

	time.Sleep(time.Second * 2)
}

func main() {
	func4()
}
