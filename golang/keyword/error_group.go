package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

/**
返回失败的错误,在这个例子中，启动了三个子任务，其中，子任务 2 会返回执行失败，其它两个执行成功。在三个子任务都执行后，group.Wait 才会返回第 2 个子任务的错误。
*/
func main() {
	var g errgroup.Group
	result := make([]error, 3)
	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})
	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		result[1] = errors.New("failed to exec #2")
		return result[1]
	})
	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")

		result[2] = errors.New("failed to exec #3")
		return result[2]
	})
	// 等待三个任务都完成
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		for _, err2 := range result {
			fmt.Println("failed:", err2)
		}
	}
}
