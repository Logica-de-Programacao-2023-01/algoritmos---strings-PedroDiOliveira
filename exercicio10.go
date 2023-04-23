package main

import "fmt"

func main() {
	palavra := "amor"
	palavra2 := "roma"

	if len(palavra) != len(palavra2) {
		fmt.Print("Não são um anagrama!")
	}
	contador := 0

	for i := 0; i < len(palavra); i++ {
		for j := 0; j < len(palavra2); j++ {
			if palavra[i] == palavra2[j] {
				contador++
			}
		}
	}
	if contador == len(palavra) {
		fmt.Println("É um anagrama!")
	} else {
		fmt.Println("Não é um anagrama!")
	}
}
