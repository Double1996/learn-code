package main

import (
	"fmt"
	"sync"
	"time"
)

//func main()  {
//	userCount := math.MaxInt64
//	for i := 0; i < userCount; i++ {
//		go func(i int) {
//			// 做一些各种各样的业务逻辑处理
//			fmt.Printf("go func: %d\n", i)
//			time.Sleep(time.Second)
//		}(i)
//	}
//}

// 尝试chan
//func main() {
//	userCount := 10
//	ch := make(chan bool, 2)
//	for i := 0; i < userCount; i++ {
//		ch <- true
//		go Read(ch, i)
//	}
//
//	time.Sleep(time.Second)
//}
//
//func Read(ch chan bool, i int) {
//	fmt.Printf("go func: %d\n", i)
//	<- ch
//}

//var wg = sync.WaitGroup{}
//
//func main() {
//	userCount := 10
//	for i := 0; i < userCount; i++ {
//		wg.Add(1)
//		go Read(i)
//	}
//
//	wg.Wait()
//}
//
//func Read(i int) {
//	defer wg.Done()
//	fmt.Printf("go func: %d\n", i)
//}

var wg = sync.WaitGroup{}

func main() {
	userCount := 10
	ch := make(chan bool, 3)
	for i := 0; i < userCount; i++ {
		go Read(ch, i)
	}
	wg.Wait()
}

func Read(ch chan bool, i int) {
	defer wg.Done() // 释放全局的锁

	wg.Add(1) // 加上一个 全局的锁
	ch <- true
	fmt.Printf("go func： %d, time:%d \n\n", i, time.Now().Unix())
	time.Sleep(time.Second)
	<-ch
}
