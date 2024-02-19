package main

import (
	"fmt"
	"log"
)

/*
Fatalはos.Exit(1)を呼び出すので、その場でプログラムが終了する
したがって、deferは実行されない
*/
func main() {
	defer fmt.Println("実行される？")
	otherFunc()
}

func otherFunc() {
	log.Fatal("something wrong")
}
