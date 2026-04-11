package main

import (
	"fmt"
	"sync"
)

type work struct {
	i int
	n int
}

func processWorker(works <-chan work, res []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for w := range works {
		res[w.i] = w.n * w.n
	}
}

func ProcessNumbers(nums []int, workers int) []int {
	if workers < 1 {
		return nil
	}

	res := make([]int, len(nums))
	workChan := make(chan work, len(nums))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go processWorker(workChan, res, &wg)
	}

	for i, n := range nums {
		workChan <- work{
			i: i,
			n: n,
		}
	}

	close(workChan)

	wg.Wait()

	return res
}

func main() {
	fmt.Println(ProcessNumbers([]int{2, 4, 6, 8}, 2))
}
