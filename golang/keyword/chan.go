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
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}

}

func main() {
	stop = make(chan bool)
	go chan1("work1")
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(3 * time.Second)
}
