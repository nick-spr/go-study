package main

import "fmt"

var fruits = map[string]int{
	"apple":  5,
	"banana": 1,
	"orange": 10,
}

func filteredFruitsByCount(minCount int) map[string]int {
	filteredFruits := make(map[string]int)

	for k, v := range fruits {
		if v > minCount {
			filteredFruits[k] = v
		}
	}

	return filteredFruits
}

func main() {
	fmt.Println(filteredFruitsByCount(2))
}
