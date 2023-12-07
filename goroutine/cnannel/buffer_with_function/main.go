package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 100)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go recvFunc(ch, wg)

	time.Sleep(3 * time.Second)

	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10

	close(ch)

	wg.Wait()
}

func recvFunc(ch <-chan int, wg *sync.WaitGroup) {
	fmt.Println("recv実行")
	// channelがcloseされるまで受信を続ける
	for i := range ch {
		fmt.Println(i)
	}
	wg.Done()
}
