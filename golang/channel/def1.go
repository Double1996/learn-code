package main

import "time"

/**
panic 只会触发goroutine 里面的延迟函数调用
*/
func main() {
	defer println("in main")
	go func() {
		defer println(" in goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)
}
