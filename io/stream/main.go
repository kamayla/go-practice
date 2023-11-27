package main

import (
	"io"
	"log"
	"os"
)

func main() {
	r, err := os.Open("io/stream/sample.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Close()

	w, err := os.Create("io/stream/write.txt")

	if err != nil {
		log.Fatalln(err)
	}

	// Readerからバイト文字列をメモリ展開せずともWriterに直接書き込める（stream処理）
	_, err = io.Copy(w, r)

	if err != nil {
		log.Fatalln(err)
	}
}
