package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		// データの加工（生産）と
		fmt.Fprint(pw, "some text")
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	// データの消費を同時に行う
	buf.ReadFrom(pr)
	fmt.Print(buf.String())
}
