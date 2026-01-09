package main

import (
	"fmt"
	"math"
)

func convert(input chan uint8, output chan float64) {
	for num := range input {
		numFloat := float64(num)
		numFloat = math.Pow(numFloat, 3)
		output <- numFloat
	}
	close(output)
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go convert(input, output)

	go func() {
		nums := []uint8{1, 2, 3, 5, 10}
		for _, v := range nums {
			input <- v
		}
		close(input)
	}()

	fmt.Println("Result:")
	for res := range output {
		fmt.Printf("%.2f\n", res)
	}

}
