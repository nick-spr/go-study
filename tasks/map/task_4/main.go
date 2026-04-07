package main

import "fmt"

var userIds = map[string]int{
	"user1": 10,
	"user2": 20,
	"user3": 30,
}

func swapMap(mapToSwap map[string]int) map[int]string {
	swappedMap := make(map[int]string)

	for k, v := range mapToSwap {
		swappedMap[v] = k
	}

	return swappedMap
}

func main() {
	fmt.Println(swapMap(userIds))
}
