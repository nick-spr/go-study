package main

import "fmt"

var testInput = []int{1, 2, 2, 3, 4, 4, 4, 5}

func countUniqueNumbers(nums []int) int {
	uniqueNumbers := make(map[int]struct{})

	for _, n := range nums {
		uniqueNumbers[n] = struct{}{}
	}

	return len(uniqueNumbers)
}

func main() {
	fmt.Println(countUniqueNumbers(testInput))
}
