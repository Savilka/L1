package main

import "fmt"

type Set struct {
	data map[string]struct{}
}

func (s Set) Add(value string) {
	if _, ok := s.data[value]; ok {
		return
	} else {
		s.data[value] = struct{}{}
	}
}

func (s Set) Remove(value string) {
	delete(s.data, value)
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := Set{map[string]struct{}{}}
	for _, s := range strings {
		set.Add(s)
	}

	for k, _ := range set.data {
		fmt.Println(k)
	}

	set.Remove("cat")

	for k, _ := range set.data {
		fmt.Println(k)
	}
}
