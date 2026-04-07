package main

import "fmt"

var a = map[string]int{
	"apple":  5,
	"banana": 2,
}

var b = map[string]int{
	"banana": 3,
	"orange": 4,
}

func sumMaps(firstMap, secondMap map[string]int) map[string]int {
	resMap := make(map[string]int)

	for k, v := range firstMap {
		resMap[k] = v
	}

	for k, v := range secondMap {
		resMap[k] += v
	}

	return resMap
}

func main() {
	fmt.Println(sumMaps(a, b))
}
