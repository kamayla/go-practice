package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}

	eg.Go(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("1")
		return nil
	})

	eg.Go(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("2")
		return nil
	})

	eg.Go(func() error {
		fmt.Println("3")
		return errors.New("3でerror")
	})

	// すべての goroutine が終わるのを待って、エラーが発生していれば（最初の）エラーを返します
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("完了")
}
