package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Сначала с помощью Sprintf получаем двоичное представление числа
// Дальше строку разбиваем на символы(биты) и меняем нужный бит
// В итоге ParseInt нам возврщает число и измененным битом

func main() {
	var number int64
	var i int
	var bit string

	fmt.Print("number: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		return
	}
	binString := fmt.Sprintf("%b", number)
	fmt.Printf("Binary %s\n", binString)

	fmt.Print("i: ")
	_, err = fmt.Scan(&i)
	if err != nil {
		return
	}

	fmt.Print("bit: ")
	_, err = fmt.Scan(&bit)
	if err != nil {
		return
	}

	splitString := strings.Split(binString, "")
	splitString[len(splitString)-i-1] = bit

	newNumber, err := strconv.ParseInt(strings.Join(splitString, ""), 2, 64)
	if err != nil {
		return
	}
	fmt.Println(newNumber)
}
