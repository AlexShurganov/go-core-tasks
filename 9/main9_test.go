package main

import (
	"slices"
	"testing"
)

func TestConvert(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go convert(input, output)

	testData := []uint8{1, 2, 3, 4}
	go func() {
		for _, v := range testData {
			input <- v
		}
		close(input)
	}()

	want := []float64{1, 8, 27, 64}
	var got []float64

	for res := range output {
		got = append(got, res)
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestConvertEmpty(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go convert(input, output)
	close(input)

	for range output {
		t.Error("unexpected data in output channel")
	}
}
