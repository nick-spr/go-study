package main

import "fmt"

var a = map[int]bool{1: true, 2: true}
var b = map[int]bool{1: true, 2: true, 3: true, 4: true}

func isSubSet(firstSet, secondSet map[int]bool) bool {
	for k := range firstSet {
		if _, ok := secondSet[k]; !ok {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isSubSet(a, b))
}
