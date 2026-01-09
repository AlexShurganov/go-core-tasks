package main

import (
	"slices"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	numChans := 3
	sources := make([]chan int, numChans)
	expectedData := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < numChans; i++ {
		sources[i] = make(chan int)
		go func(ch chan int, start int) {
			defer close(ch)
			ch <- start
			ch <- start + 1
		}(sources[i], (i*2)+1)
	}

	dest := merge(sources)

	var got []int
	for val := range dest {
		got = append(got, val)
	}

	sort.Ints(got)

	if !slices.Equal(got, expectedData) {
		t.Errorf("got %v, want %v", got, expectedData)
	}
}

func TestMergeEmpty(t *testing.T) {
	var sources []chan int
	dest := merge(sources)

	_, ok := <-dest
	if ok {
		t.Error("expected closed channel for empty input")
	}
}

func TestMergeAllClosed(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	close(ch1)
	close(ch2)

	dest := merge([]chan int{ch1, ch2})

	count := 0
	for range dest {
		count++
	}

	if count != 0 {
		t.Errorf("got %d elements, want 0", count)
	}
}
