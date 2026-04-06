package main

import (
	"fmt"
)

type Account struct {
	Id   int64
	Name string
}

var accounts = []Account{
	{
		Id:   1,
		Name: "first",
	},
	{
		Id:   2,
		Name: "second",
	},
	{
		Id:   3,
		Name: "third",
	},
	{
		Id:   4,
		Name: "fourth",
	},
}

func GetAccountsPointers() []*Account {
	res := make([]*Account, 0, len(accounts))
	for _, a := range accounts {
		res = append(res, &a)
	}
	return res
}

func main() {
	fmt.Println("---------- Вопрос №1 ----------")
	// Что вывелось в консоль?
	account := Account{Id: 23, Name: "Mike"}
	fmt.Printf("%p \n", &account) // адрес памяти - указатель на переменную account 0x...

	fmt.Println("---------- Вопрос №2 ----------")
	// Что вывелось в консоль?
	// Почему первый первый и второй вывод отличаются?
	account2 := account
	fmt.Printf("%p \n", &account2) // адрес памяти - указатель на переменную account2 - 0x...
	fmt.Printf("%p \n", &account)  // адрес памяти - указатель на переменную account - 0x...
	// адреса отличаются, поскольку это две разные переменные, которые хранятся в разных участках памяти
	// когда сделали account2 := account мы скопировали структуру из account в новую переменную account2

	fmt.Println("---------- Вопрос №3 ----------")
	// Какие значения будут храниться в переменных a b c?
	// Объясните каждый из выводов в консоль
	a := 2
	b := &a
	c := &a
	a++
	fmt.Println(&a, &b, &c) // адрес памяти - указатель на переменную a - 0x..., указатель на указатель на переменную а, указатель на указатель на переменную а (отличный от предыдущего, поскольку переменная, которая его хранит - другая)
	fmt.Println(a, b, c)    // 3, адрес памяти - указатель на переменную a - 0x..., адрес памяти - указатель на переменную a - 0x...
	fmt.Println(a, *b, *c)  // 3, 3, 3 - поскольку переменные b и с - указатели на a - получаем новое значение (3, а не 2) при их разименовывании

	fmt.Println("---------- Вопрос №4 ----------")
	// Почему в консоль постоянно выводятся различне адреса памяти,
	// хотя при каждом вызове функции GetAccountsPointers мы работаем с одним и тем же accounts?

	// Вызов a1
	pointers := GetAccountsPointers()
	for _, pointer := range pointers {
		fmt.Printf("%p \n", pointer)
	}
	fmt.Println()
	// Вызов a2
	pointers = GetAccountsPointers()
	for _, pointer := range pointers {
		fmt.Printf("%p \n", pointer)
	}
	// внутри GetAccountsPointers мы перекладывваем элементы из слайса accounts в новый слайс через промежуточную переменную a и в слайс кладём именно указатель на неё
	// чтобы избежать этого, необходимо использовать индекс, а не сохранять элементы в промежуточную переменную

	fmt.Println("---------- Вопрос №5 ----------")
	// Почему в sliceOfSlice 9 вывелась, а в copyOfS нет?
	// в цикле for в переменную s копируется массив [2]int и слайс делается на этой локальной копии, а не на изначальном массиве
	// локальная копия s ничего не знает про изменения в начальном массиве, а слайс ссылается на неё

	// почему используется s[:], а не s?
	// тип значений sliceOfSlice - [][2]int, соответственно при range мы получаем массив с типом [2]int
	// тип copyOfS - [][]int - слайс слайсов - в него нельзя положить массив [2]int - поэтому делаем слайс всего массива

	// если у sliceOfSlice и copyOfS будет тип [][]int, как сделать чтобы 9 была только в sliceOfSlice?
	// использовать функцию copy для принудительного копирования данных в новый участов памяти
	// 	temp := make([]int, len(s))
	//  copy(temp, s)
	//  copyOfS = append(copyOfS, temp)

	sliceOfSlice := [][2]int{{1, 2}, {3, 4}, {5, 6}}
	var copyOfS [][]int
	for _, s := range sliceOfSlice {
		copyOfS = append(copyOfS, s[:])
	}
	sliceOfSlice[0][0] = 9
	fmt.Println(sliceOfSlice)
	fmt.Println(copyOfS)
}
