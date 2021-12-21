package main

import "fmt"

// Producer
// https://cloud.tencent.com/developer/article/1559262
// 通过chan在生产者和消费者之间传递数据(ch)和同步状态(done)；
// chan作为参数传递时是引用传递，不需要使用指针；
// chan是协程安全的，多个goroutine之间不需要锁；
// chan的close事件可以被recv获取，close事件一定在正常数据读完之后，机制类似于read到EOF;
/**

生产者：发送数据端

消费者：接收数据端

缓冲区：
　　1. 解耦(降低生产者和消费者之间耦合度)

　　2. 并发(生产者消费者数量不对等时，能保持正常通信)

　　3. 缓存(生产者和消费者 数据处理速度不一致时,暂存数据)


1：解耦
假设生产者和消费者分别是两个类。如果让生产者直接调用消费者的某个方法，那么生产者对于消费者就会产生依赖（也就是耦合）。将来如果消费者的代码发生变化，可能会直接影响到生产者。而如果两者都依赖于某个缓冲区，两者之间不直接依赖，耦合度也就相应降低了。


2：处理并发
生产者直接调用消费者的某个方法，还有另一个弊端。由于函数调用是同步的（或者叫阻塞的），在消费者的方法没有返回之前，生产者只好一直等在那边。万一消费者处理数据很慢，生产者只能无端浪费时间。
使用了生产者／消费者模式之后，生产者和消费者可以是两个独立的并发主体。生产者把制造出来的数据往缓冲区一丢，就可以再去生产下一个数据。基本上不用依赖消费者的处理速度。
其实最当初这个生产者消费者模式，主要就是用来处理并发问题的。

3：缓存
如果生产者制造数据的速度时快时慢，缓冲区的好处就体现出来了。当数据制造快的时候，消费者来不及处理，未处理的数据可以暂时存在缓冲区中。等生产者的制造速度慢下来，消费者再慢慢处理掉。

*/
///**
func Producer(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	close(ch) // 不去close channel 会deadlock
}

func Consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("id: %d, recv: %d\n", id, value)
		} else {
			fmt.Printf("id: %d, closed:\n", id)
			break
		}
	}
	done <- true
}

func main() {
	ch := make(chan int, 3)
	coNum := 2
	done := make(chan bool, coNum)
	for i := 1; i <= coNum; i++ { // 两个消费者
		go Consumer(i, ch, done) // 然后开启一个新的goroutine，把双向通道作为参数传递到producer方法中，同时转成只写通道。子协程开始执行循环，向只写通道中添加数据，这就是生产者。
	}
	go Producer(ch) // 主协程，直接调用consumer方法，该方法将双向通道转成只读通道，通过循环每次从通道中读取数据，这就是消费者。 10 个消费者
	for i := 0; i < coNum; i++ {
		<-done
	}
}
