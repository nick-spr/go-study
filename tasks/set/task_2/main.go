package main

import "fmt"

var a = map[int]bool{10: true, 20: true, 30: true, 40: true}
var b = map[int]bool{20: true, 40: true}

func setDiff(firstSet, secondSet map[int]bool) map[int]bool {
	res := make(map[int]bool)

	for k, v := range firstSet {
		if _, ok := secondSet[k]; !ok {
			res[k] = v
		}
	}

	return res
}

func main() {
	fmt.Println(setDiff(a, b))
}
