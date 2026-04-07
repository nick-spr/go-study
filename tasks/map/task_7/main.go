package main

import "fmt"

var featureFlags = map[string]bool{
	"featureA": true,
	"featureB": false,
	"featureC": true,
}

func invertFlags(featureMap map[string]bool) map[string]bool {
	resMap := make(map[string]bool)

	for k, v := range featureMap {
		resMap[k] = !v
	}

	return resMap
}

func main() {
	fmt.Println(invertFlags(featureFlags))
}
