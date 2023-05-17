package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const IncrementTimes = 20000

func AtomicIncrement(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < IncrementTimes; i++ {
		atomic.AddInt64(counter, 1)
	}
}

func MutexIncrement(counter *int64, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	// 这样处理会比每次循环都拿锁要快很多，也会比原子操作快
	mutex.Lock()
	for i := 0; i < IncrementTimes; i++ {
		*counter++
	}
	mutex.Unlock()
}

func ConcurrentAtomicAdd() int64 {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var counter int64 = 0
	go AtomicIncrement(&counter, &wg)
	go AtomicIncrement(&counter, &wg)
	wg.Wait()
	return counter
}
func ConcurrentMutexAdd() int64 {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(2)
	var counter int64 = 0
	go MutexIncrement(&counter, &wg, &mutex)
	go MutexIncrement(&counter, &wg, &mutex)
	wg.Wait()
	return counter
}

func main() {
	fmt.Printf("MutexAdd result: %d\n", ConcurrentMutexAdd())
	fmt.Printf("AtomicAdd result: %d\n", ConcurrentAtomicAdd())
}
