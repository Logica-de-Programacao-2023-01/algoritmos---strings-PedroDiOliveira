package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma string em letra minúscula e sem acentos ")
	scanner.Scan()
	frase1 := scanner.Text()
	frase1 = strings.ReplaceAll(frase1, "a", "")
	frase1 = strings.ReplaceAll(frase1, "e", "")
	frase1 = strings.ReplaceAll(frase1, "i", "")
	frase1 = strings.ReplaceAll(frase1, "o", "")
	frase1 = strings.ReplaceAll(frase1, "u", "")
	fmt.Println(frase1)
}

//CAMINHO MAIS COMPLEXO
//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func removeVowels(s string) string {
//	// Define as vogais que serão removidas
//	vowels := "aeiouAEIOU"
//
//	// Cria uma função que verifica se um caractere é uma vogal
//	isVowel := func(char rune) bool {
//		return strings.IndexRune(vowels, char) >= 0
//	}
//
//	// Remove as vogais da string utilizando a função Filter do pacote strings
//	result := strings.Map(func(char rune) rune {
//		if isVowel(char) {
//			return -1
//		}
//		return char
//	}, s)
//
//	return result
//}
//
//func main() {
//	// Lê a string do usuário
//	var input string
//	fmt.Print("Digite uma string: ")
//	fmt.Scanln(&input)
//
//	// Remove as vogais da string
//	result := removeVowels(input)
//
//	// Imprime o resultado
//	fmt.Println("Resultado:", result)
//}
