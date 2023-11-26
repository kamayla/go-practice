package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func main() {
	result := doSomeThingParallel(10)
	log.Println(result)
}

func doSomeThingParallel(workerNum int) error {
	ctx := context.Background()
	// ctxからcancel可能なctxを作成
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	errCh := make(chan error, workerNum)

	wg := sync.WaitGroup{}

	for i := 0; i < workerNum; i++ {
		i := i
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if err := doSomeThingWithContext(cancelCtx, num); err != nil {
				cancel()
				errCh <- err
			}
		}(i)
	}

	// wg.Addで追加した数分Doneされるまで処理を待つ
	wg.Wait()

	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func doSomeThingWithContext(ctx context.Context, num int) error {
	select {
	// 処理に入る前にコンテキストの死活確認を行う（死んでたらエラーを返して処理をしない）
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	fmt.Println(num)
	return nil
}
