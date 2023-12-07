package main

import (
	"fmt"
	"time"
)

func main() {
	// バッファなしchannelを作成
	ch := make(chan int)

	go func() {
		fmt.Println("実行")
		time.Sleep(5 * time.Second)
		ch <- 10
		ch <- 10
		ch <- 10
		ch <- 10
		ch <- 10
		fmt.Println("バッファなしは、受信側が受け取るまで送信側はブロックされる")
	}()

	<-ch

	fmt.Println("処理がきた")
}
