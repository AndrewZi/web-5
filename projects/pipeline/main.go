package main

import (
	"fmt"
)

func removeDuplicates(in, out chan string) {
	var buf []string
	for i := range in {
		flag := false
		for _, j := range buf {
			if i == j {
				flag = true
			}
		}
		if flag == false {
			out <- i
		}
		buf = append(buf, i)
	}
	close(out)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)
		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
}
