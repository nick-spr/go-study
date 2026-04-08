package main

import "fmt"

var a = map[string]bool{"go": true, "java": true, "rust": true}
var b = map[string]bool{"rust": true, "python": true}

func combineSets(firstSet, secondSet map[string]bool) map[string]bool {
	res := make(map[string]bool)

	for k, v := range firstSet {
		res[k] = v
	}

	for k, v := range secondSet {
		res[k] = v
	}

	return res
}

func main() {
	fmt.Println(combineSets(a, b))
}
