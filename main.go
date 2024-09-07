package main

import "fmt"

const prefixoOlaPortugues string = "Olá, "

func main() {
	fmt.Println(Ola("nome"))
}

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoOlaPortugues + nome
}
