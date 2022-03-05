package main


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
