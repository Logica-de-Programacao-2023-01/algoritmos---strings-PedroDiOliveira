package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, char, replace string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	fmt.Print("Digite o caractere que você deseja substituir: ")
	fmt.Scanln(&char)

	fmt.Print("Digite o caractere pelo qual você deseja substituir: ")
	fmt.Scanln(&replace)

	newStr := strings.ReplaceAll(str, char, replace)

	fmt.Println("Nova string: ", newStr)
}
