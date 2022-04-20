package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Родительская структура
type Human struct {
	age int
}

// Метод родительской структуры
func (h Human) getAge() int {
	return h.age
}

// Структура с встренной структурой Human. Наследуются все методы и поля
type Action struct {
	Human
}

// Метод можно переопределять
//func (a Action) getAge() int {
//	return a.Age + 10
//}
// a.getAge() : 25
// a.Human.getAge() : 15

func main() {
	h := Human{age: 15}
	a := Action{Human: h}
	fmt.Println(a.getAge())
	fmt.Println(a.age)
}
