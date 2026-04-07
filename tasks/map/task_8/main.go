package main

import "fmt"

var ratings = map[int]int{
	1: 10,
	2: 3,
	3: 7,
	4: 1,
}

func deleteByRating(ratingsMap map[int]int, rating int) {
	for k, v := range ratingsMap {
		if v < rating {
			delete(ratingsMap, k)
		}
	}
}

func main() {
	deleteByRating(ratings, 5)
	fmt.Println(ratings)
}
