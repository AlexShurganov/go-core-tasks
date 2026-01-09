package main

import (
	"fmt"
	"math/rand/v2"
)

func sliceExample(slice []int) []int {
	var res []int

	for _, v := range slice {
		if v%2 == 0 {
			res = append(res, v)
		}
	}

	return res
}

func addElements(slice []int, number int) []int {
	slice = append(slice, number)
	return slice
}

func removeElement(slice []int, index int) []int {
	res := make([]int, len(slice)-1)
	copy(res[:index], slice[:index])
	copy(res[index:], slice[index+1:])
	return res
}

func copySlice(slice []int) []int {
	res := make([]int, len(slice))
	copy(res, slice)
	return res
}

func generateSlice() []int {
	res := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		res = append(res, rand.IntN(100))
	}
	return res
}

func main() {
	originalSlice := generateSlice()
	fmt.Println("Created slice: ", originalSlice)

	evens := sliceExample(originalSlice)
	fmt.Println("Only evens: ", evens)

	appendedSlice := addElements(originalSlice, 7)
	fmt.Println("Slice after adding 7: ", appendedSlice)

	removedSlice := removeElement(originalSlice, 1)
	fmt.Println("Deleted index 1: ", removedSlice)

	newSlice := copySlice(originalSlice)
	fmt.Println("Slice after copying: ", newSlice)

}
