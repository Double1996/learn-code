package main

import (
	"context"
	"fmt"
	"time"
)

// 控制单个协程
func context1(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // 调用 cancel() 协程进行取消操作
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

type Options struct{ Interval time.Duration }

func context2(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // 调用 cancel() 协程进行取消操作
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

func context3(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // 调用 cancel() 协程进行取消操作
			fmt.Println("stop", name, ctx.Err())
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 1.创建可以取消的协程
	//ctx, cancel := context.WithCancel(context.Background())
	//go context1(ctx, "worker1")
	//time.Sleep(3 * time.Second)
	//cancel()
	//time.Sleep(3 * time.Second)

	// 2.控制多个协程
	//ctx, cancel := context.WithCancel(context.Background())
	//go context1(ctx, "worker1")
	//go context1(ctx, "worker2")
	//time.Sleep(3* time.Second)
	//cancel()
	//time.Sleep(3* time.Second)

	// 3.子协程 如果需要往子协程中传递参数，使用
	//ctx, cancel := context.WithCancel(context.Background())
	//vCtx := context.WithValue(ctx, "options", &Options{1})
	//
	//go context2(vCtx, "worker1")
	//go context2(vCtx, "worker2")
	//
	//time.Sleep(3 * time.Second)
	//cancel()
	//time.Sleep(3 * time.Second)

	// 4.context WithTimeout,  设置 超时控制时间
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//go context1(ctx, "worker1")
	//go context1(ctx, "worker2")
	//
	//time.Sleep(3 * time.Second)
	//fmt.Println("before cancel")
	//cancel()
	//time.Sleep(3 * time.Second)

	// 5. deadline 控制最少的退出时间 deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	go context3(ctx, "worker1")
	go context3(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}
