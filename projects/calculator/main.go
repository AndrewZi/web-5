package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resChan := make(chan int)
	go func() {
		defer close(resChan)
		select {
		case val := <-firstChan:
			resChan <- val * val
		case val := <-secondChan:
			resChan <- val * 3
		case <-stopChan:
			return
		}
	}()
	return resChan
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan struct{})
	res := calculator(ch1, ch2, ch3)

	ch2 <- 16

	fmt.Println(<-res)
}
