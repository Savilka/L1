package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.Replace(str, "\r\n", "", -1)
	words := strings.Split(str, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	str = strings.Join(words, " ")
	fmt.Println(str)
}
