package interview

import (
	"fmt"
	"sync"
	"unicode/utf8"
)

/**
  使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
  12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

/**
  // 思路： 通过 chan 信号量来进行控制
  //
  //
*/

func PrintWordNumber() {
	wordCh, numberCh := make(chan struct{}), make(chan struct{}) // 无缓冲的通道
	wg := sync.WaitGroup{}                                       // 控制同步

	go func() {
		i := 1
		for {
			<-numberCh
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
			if i >= utf8.RuneCountInString(str) { // 推出
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
