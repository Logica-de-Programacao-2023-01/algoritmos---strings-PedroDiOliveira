package main

import (
	"fmt"
	"unicode"
)

// Solicite ao usuário uma string e informe se ela contém pelo menos um número.
func main() {
	//variaveis
	var str string

	//pede a string ao usuario
	fmt.Println("Informe a string:")
	fmt.Scan(&str)

	contemNumero := false

	for _, r := range str {
		if unicode.IsDigit(r) {
			contemNumero = true
			break
		}
	}
	if contemNumero {
		fmt.Println("A string contem pelo menos um numero!")
	} else {
		fmt.Println("A string não contem nenhum numero!")
	}
}
