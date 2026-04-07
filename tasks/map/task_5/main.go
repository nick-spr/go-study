package main

import (
	"fmt"
	"unicode/utf8"
)

var testSlice = []string{"go", "java", "rust", "c", "js", "Б"}

func groupStrings(stringsToGroup []string) map[int][]string {
	stringLenMap := make(map[int][]string)

	for _, s := range stringsToGroup {
		stringLenMap[utf8.RuneCountInString(s)] = append(stringLenMap[utf8.RuneCountInString(s)], s)
		// если уверены, что будут только ASCII символы, то можно использовать len(s)
	}

	return stringLenMap
}

func main() {
	fmt.Println(groupStrings(testSlice))
}
