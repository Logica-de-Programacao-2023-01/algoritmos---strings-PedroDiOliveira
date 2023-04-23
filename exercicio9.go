package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.
func main() {
	//variaveis
	var x, a, b string

	//Entrada de dados
	fmt.Println("Informe a string:")
	fmt.Scan(&x)
	fmt.Println("Informe o caracter a ser substituido:")
	fmt.Scan(&a)
	fmt.Println("Informe o que deve substituilo:")
	fmt.Scan(&b)

	//Saida de dados
	resultado := strings.ReplaceAll(x, a, b)
	fmt.Println(resultado)
}
