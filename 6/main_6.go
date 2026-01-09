/*

Напишите генератор случайных чисел используя небуфферизированный канал.

* Напишите unit тесты к созданным функциям

*/

package main

import (
	"fmt"
	"math/rand/v2"
)

func GenerateRandomNumber(ch chan int, done chan struct{}) {
	for {
		select {
		case <-done:
			close(ch)
			return
		case ch <- rand.IntN(1000):
		}
	}
}

func main() {
	randomNumbers := make(chan int)
	done := make(chan struct{})

	go GenerateRandomNumber(randomNumbers, done)

	count := 0
	for v := range randomNumbers {
		fmt.Println(v)
		count++
		if count >= 10 {
			close(done)
		}
	}
}
