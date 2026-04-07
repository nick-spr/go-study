package main

import "fmt"

var testStrings = []string{"go", "java", "go", "python", "go", "java"}

func countWords(words []string) map[string]int {
	var wordsCount = make(map[string]int)

	for _, word := range words {
		wordsCount[word] += 1
	}

	return wordsCount
}

func main() {
	fmt.Println(countWords(testStrings))
}
