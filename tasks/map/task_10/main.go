package main

import "fmt"

var testInput = "hello world"

func runeCounter(text string) map[rune]int {
	res := make(map[rune]int)

	for _, r := range text {
		if r == ' ' { // пробелы не учитываем
			continue
		}
		res[r]++
	}

	return res
}

func main() {
	runeCount := runeCounter(testInput)
	fmt.Println(runeCount)

	for k, v := range runeCount {
		fmt.Println(string(k), ":", v)
	}
}
