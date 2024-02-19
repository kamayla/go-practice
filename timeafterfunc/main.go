package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

// 3秒ごとに繰り返し処理みたいなことができる
func main() {
	var timer *time.Timer
	timer = time.AfterFunc(3*time.Second, func() {
		println("Hello, world!")
		timer.Reset(3 * time.Second)
	})

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, os.Interrupt)

	log.Println(<-sigCh)
	timer.Stop()
}
