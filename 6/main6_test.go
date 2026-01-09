package main

import (
	"testing"
	"time"
)

func TestGenerateRandomNumber(t *testing.T) {
	ch := make(chan int)
	done := make(chan struct{})

	go GenerateRandomNumber(ch, done)

	select {
	case val := <-ch:
		if val < 0 || val >= 1000 {
			t.Errorf("got number %d, want range [0, 1000)", val)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timeout")
	}

	close(done)

	count := 0
	for range ch {
		count++
		if count > 100 {
			t.Fatal("channel leak")
		}
	}
}

func TestGeneratorUniqueness(t *testing.T) {
	ch := make(chan int)
	done := make(chan struct{})
	defer close(done)

	go GenerateRandomNumber(ch, done)

	n1 := <-ch
	n2 := <-ch

	if n1 == n2 {
		t.Logf("identical numbers received: %d and %d", n1, n2)
	}
}
