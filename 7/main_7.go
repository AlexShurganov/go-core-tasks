package main

import (
	"fmt"
	"sync"
)

func merge(sources []chan int) chan int {
	dest := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(sources))

	for _, ch := range sources {
		go func(ch chan int) {
			defer wg.Done()
			for val := range ch {
				dest <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest

}

func main() {

	channels := make([]chan int, 5)

	for i := range channels {
		channel := make(chan int)
		channels[i] = channel

		go func(ch chan int) {
			defer close(ch)
			for j := 1; j <= 5; j++ {
				ch <- j
			}
		}(channel)
	}

	dest := merge(channels)

	for val := range dest {
		fmt.Printf("Received: %d\n", val)
	}

}
