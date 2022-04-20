package main

import (
	"encoding/json"
	"fmt"
)

// Интерфейс, к которому адаптируем структуру JsonP
type Printer interface {
	Print()
}

// Какая-то структура уже реализующая интерфейс
type StringPrinter struct {
	Data string
}

func (p StringPrinter) Print() {
	fmt.Println(p.Data)
}

// Структура, которую адаптируем
type JsonP struct {
	Data struct {
		Data string `json:"data"`
	}
}

// Метод не подоходящий под сигнатуру интерфейса
func (j JsonP) JsonPrint() {
	str, err := json.MarshalIndent(j.Data, "", " ")
	if err != nil {
		return
	}
	fmt.Println(string(str))
}

// Адаптер наследует все поля и методы адаптируемой структуры
type Adapter struct {
	*JsonP
}

// Метод подоходящий под сигнатуру интерфейса
func (a Adapter) Print() {
	a.JsonPrint()
}

func main() {
	sp := StringPrinter{Data: "DATA"}
	jp := JsonP{Data: struct {
		Data string `json:"data"`
	}(struct{ Data string }{Data: "DATA"})}
	ad := Adapter{JsonP: &jp}
	sp.Print()
	ad.Print()
}
