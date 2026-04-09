package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		atomic_go()
	}()

	go func() {
		defer wg.Done()
		mutex_go()
	}()

	go func() {
		defer wg.Done()
		chan_go()
	}()

	wg.Wait()
	fmt.Println("Результаты готовы!")
}

func atomic_go() {
	counter := int64(0)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Результат атомиков:", counter)
}

func mutex_go() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock()

			mu.Lock()
			counter = counter + 1
		}()
	}

	wg.Wait()
	fmt.Println("Результат мьютексов:", counter)
}

func chan_go() {
	counter := 0
	ch := make(chan int)

	for i := 0; i < 100; i++ {
		go func() {
			ch <- 1
		}()
	}

	for i := 0; i < 100; i++ {
		counter += <-ch
	}

	fmt.Println("Результат каналов:", counter)
}
