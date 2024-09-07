package main

const espanhol string = "espanhol"
const frances string = "frances"
const prefixoOlaPortugues string = "Olá, "
const prefixoOlaEspanhol string = "Hola, "
const prefixoOlaFrances string = "Bonjour, "

func main() {
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	if idioma == espanhol {
		return prefixoOlaEspanhol + nome
	}

	if idioma == frances {
		return prefixoOlaFrances + nome
	}

	return prefixoOlaPortugues + nome
}
