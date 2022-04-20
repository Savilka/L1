package main

import (
	"fmt"
	"strings"
)

// Сначала переводим строку в нижний регистр. Далее создаём map, где будем хранить информацию о вхождении каждого символа.
// Тип struct{} был выбран для экономии памяти

func checkUniqSymbols(str string) bool {
	str = strings.ToLower(str)
	marks := make(map[rune]struct{})
	for _, symbol := range str {
		_, ok := marks[symbol]
		if ok {
			return false
		}
		marks[symbol] = struct{}{}
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(checkUniqSymbols(str))

}
