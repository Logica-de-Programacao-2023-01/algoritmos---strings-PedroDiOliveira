package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	//perguntando ao usuario
	fmt.Println("Qual a string?")
	fmt.Scan(&str)
	conversao := strings.ToUpper(str)
	fmt.Println("A sua string convertida para maiusculo fica:", conversao)
}
