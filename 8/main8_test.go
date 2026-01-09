package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestSemaphoreWG_Basic(t *testing.T) {
	size := 3
	swg := NewSemaphoreWG(size)

	var counter atomic.Int32

	swg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer swg.Done()
			time.Sleep(50 * time.Millisecond)
			counter.Add(1)
		}()
	}

	swg.Wait()

	if counter.Load() != 3 {
		t.Errorf("Wait finished prematurely: got %d tasks completed, want 3", counter.Load())
	}
}

func TestSemaphoreWG_WaitBlocks(t *testing.T) {
	swg := NewSemaphoreWG(1)
	swg.Add(1)

	waitFinished := make(chan bool)

	go func() {
		swg.Wait()
		waitFinished <- true
	}()

	select {
	case <-waitFinished:
		t.Fatal("Wait did not block execution, although Done was not called")
	case <-time.After(100 * time.Millisecond):

	}

	swg.Done()

	select {
	case <-waitFinished:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Wait did not unlock after Done")
	}
}

func TestSemaphoreWG_CapLimit(t *testing.T) {
	size := 5
	swg := NewSemaphoreWG(size)

	if cap(swg.sem) != size {
		t.Errorf("channel capacity = %d; want %d", cap(swg.sem), size)
	}
}
