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

func main() {
	//stop = make(chan bool)
	//go chan1("work1")
	//time.Sleep(3 * time.Second)
	//stop <- true
	//time.Sleep(3 * time.Second)
	chan4()
}
