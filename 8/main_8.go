/*
Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.

* Напишите unit тесты к созданным функциям
*/

package main

type SemaphoreWG struct {
	sem chan struct{}
}

func NewSemaphoreWG(size int) *SemaphoreWG {
	return &SemaphoreWG{
		sem: make(chan struct{}, size),
	}
}

func (swg *SemaphoreWG) Add(n int) {
	for i := 0; i < n; i++ {
		swg.sem <- struct{}{}
	}
}

func (swg *SemaphoreWG) Done() {
	<-swg.sem
}

func (swg *SemaphoreWG) Wait() {
	limit := cap(swg.sem)
	for i := 0; i < limit; i++ {
		swg.sem <- struct{}{}
	}

	for i := 0; i < limit; i++ {
		<-swg.sem
	}
}
