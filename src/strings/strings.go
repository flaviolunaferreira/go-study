package main

import "fmt"

func main() {
	fmt.Println("Hello Word") //Imprimindo texto na tela
	fmt.Println("quantidade de caracteres -> ", len("Hello Word")) // Pegando a quantidade de caracteres
	fmt.Println("número binário da letra -> ", "Hello Word"[2]) // Pegando caracter específico -> index inicia com 0
	fmt.Println("Hello" + " Word") // concatenando strings
}