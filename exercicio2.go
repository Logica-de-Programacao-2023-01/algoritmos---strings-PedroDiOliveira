package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Digite uma string:")
	fmt.Scan(&str)

	newStr := strings.Join(strings.Fields(str), " ")

	fmt.Printf("A nova string é: %s\n", newStr)
}
