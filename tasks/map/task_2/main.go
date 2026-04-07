package main

import "fmt"

var users = map[string]bool{
	"alice": true,
	"bob":   false,
}

func isUserExist(user string) bool {
	_, ok := users[user]
	return ok
}

func main() {
	fmt.Println(isUserExist("alice"))
	fmt.Println(isUserExist("bob"))
	fmt.Println(isUserExist("charlie"))
}
