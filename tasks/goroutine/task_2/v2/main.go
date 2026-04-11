package main

import (
	"fmt"
	"sync"
)

type work struct {
	i int
	n int
}

func processWorker(works <-chan work, res chan<- work, wg *sync.WaitGroup) {
	defer wg.Done()
	for w := range works {
		res <- work{
			i: w.i,
			n: w.n * w.n,
		}
	}
}

func ProcessNumbers(nums []int, workers int) []int {
	if workers < 1 {
		return nil
	}

	res := make([]int, len(nums))
	workChan := make(chan work, workers)
	resChan := make(chan work, workers)
	var workersWg sync.WaitGroup

	for i := 0; i < workers; i++ {
		workersWg.Add(1)
		go processWorker(workChan, resChan, &workersWg)
	}

	go func() {
		for i, n := range nums {
			workChan <- work{
				i: i,
				n: n,
			}
		}
		close(workChan)
	}()

	go func() {
		workersWg.Wait()
		close(resChan)
	}()

	for w := range resChan {
		res[w.i] = w.n
	}

	return res
}

func main() {
	fmt.Println(ProcessNumbers([]int{2, 4, 6, 8}, 3))
}
