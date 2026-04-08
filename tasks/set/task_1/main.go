package main

import "fmt"

var a = map[int]bool{1: true, 2: true, 3: true, 5: true}
var b = map[int]bool{3: true, 4: true, 5: true, 6: true}

func repeatingElements(firstSet, secondSet map[int]bool) map[int]bool {
	res := make(map[int]bool)

	for k, v := range firstSet {
		if _, ok := secondSet[k]; ok {
			res[k] = v
		}
	}

	return res
}

func main() {
	fmt.Println(repeatingElements(a, b))
}
