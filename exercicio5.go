package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("Digite um numero em ponto flutuante:")
	fmt.Scan(&str)

	_, err := strconv.ParseFloat(str, 64)

	if err == nil {
		fmt.Println("É um numero de ponto flutuante valido")
	} else {
		fmt.Println("Não é um numero de ponto flutuante valido")
	}
}
