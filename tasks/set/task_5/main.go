package main

import "fmt"

var a = map[int]bool{1: true, 2: true, 3: true}
var b = map[int]bool{3: true, 4: true, 5: true}

func symmetricalSetDiff(firstSet, secondSet map[int]bool) map[int]bool {
	res := make(map[int]bool)

	for k, v := range secondSet {
		res[k] = v
	}

	for k, v := range firstSet {
		if _, ok := res[k]; !ok {
			res[k] = v
		} else {
			delete(res, k)
		}
	}

	return res
}

func main() {
	fmt.Println(symmetricalSetDiff(a, b))
}
