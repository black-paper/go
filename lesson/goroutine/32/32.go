package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	// 何か処理を実行
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			// 処理が成功しても失敗してもwg.Done()が実行されるようにInnerFunctionで実行するのはあり
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer生成
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer生成
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
}
