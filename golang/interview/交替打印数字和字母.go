package interview

import (
	"fmt"
	"sync"
)

/**
  使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
  12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

/**
  // 思路： 通过 chan 信号量来进行控制
  // 掌握 基于 信号通知的写法， 几个要点
  // 1. sync.waitGroup 阻塞
  // 2. 定义不同的 channel 进行 按顺序消息传递
  // 3. 等待满足结果退出
*/

func PrintWordNumber() {
	wordCh, numberCh := make(chan struct{}), make(chan struct{}) // 无缓冲的通道
	wg := sync.WaitGroup{}                                       // 控制同步

	go func() {
		i := 1
		for {
			<-numberCh // 开始执行
			fmt.Printf("%d%d", i, i+1)
			i += 2
			wordCh <- struct{}{} // 进行信号通知
		}
	}()
	wg.Add(1)

	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			<-wordCh
			if i >= len(str) { // 同步成功
				wait.Done()
				return
			}
			fmt.Print(str[i : i+2])
			i += 2
			numberCh <- struct{}{}
		}
	}(&wg)

	numberCh <- struct{}{} // 依次信号量通知
	wg.Wait()
}
