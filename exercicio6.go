package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&input)

	// Remove espaços em branco do começo e do final da string
	input = strings.TrimSpace(input)

	// Conta quantas palavras existem na string
	words := strings.Split(input, " ")
	count := len(words)

	fmt.Printf("A frase digitada contém %d palavras.\n", count)
}
