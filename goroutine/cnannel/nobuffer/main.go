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
	}()

	// バッファなしチャネルは受信があるまで処理をブロッキングする
	<-ch

	fmt.Println("処理がきた")
}
