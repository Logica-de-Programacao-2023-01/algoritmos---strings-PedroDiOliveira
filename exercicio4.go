package main

import "fmt"

// Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.
func main() {
	//Variaveis
	var str string
	var str2 string

	//preenchendo as variaveis
	fmt.Println("Qual é a primeira string?")
	fmt.Scan(&str)
	fmt.Println("Qual é a segunda string?")
	fmt.Scan(&str2)

	if str == str2 {
		fmt.Println("As strings são iguais!")
	} else {
		fmt.Println("As strings são diferentes!")
	}
}
