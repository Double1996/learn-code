package main

import (
	"fmt"
	"time"
)

var stop chan bool

// 简单的生产者消费者模型
func chan1(name string) {
	for {
		select {
		case <-stop:
			fmt.Println("stop", name) // 协程主动的退出
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

// 无缓冲的
func chan2() {

}

// 有缓冲的通道
func chan3() {

}

// 读取关闭的 chan
func chan4() {
	ch := make(chan int, 5)
	ch <- 18
	close(ch)
	x, ok := <-ch
	if ok {
		fmt.Println("received: ", x)
	}

	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.")
	}
}

// chan 的泄露问题
func process(timeout time.Duration) bool {
	ch := make(chan bool)
	go func() {
		// 模拟处理耗时的业务
		time.Sleep(timeout + time.Second)
		ch <- true // block
		fmt.Println("exit goroutine")
	}()
	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return false
	}
}

func main() {
	//stop = make(chan bool)
	//go chan1("work1")
	//time.Sleep(3 * time.Second)
	//stop <- true
	//time.Sleep(3 * time.Second)
	//chan4()
	process(5 * time.Second)
}
