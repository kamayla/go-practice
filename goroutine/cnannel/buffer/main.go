package main

import (
	"fmt"
	"time"
)

func main() {
	// バッファありchannelを作成
	ch := make(chan int, 100)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 10
		ch <- 10
		ch <- 10
		ch <- 10
		ch <- 10
		ch <- 10
		ch <- 10
		ch <- 10
		fmt.Println("バッファありは、バッファがいっぱいになるまで送信側はブロックされない")
	}()

	<-ch

	fmt.Println("終了")
}
