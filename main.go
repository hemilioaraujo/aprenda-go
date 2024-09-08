package main

const alemao string = "alemao"
const espanhol string = "espanhol"
const frances string = "frances"
const prefixoOlaAlemao string = "Hallo, "
const prefixoOlaEspanhol string = "Hola, "
const prefixoOlaFrances string = "Bonjour, "
const prefixoOlaPortugues string = "Olá, "

func main() {
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	prefixo := prefixoOlaPortugues

	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case alemao:
		prefixo = prefixoOlaAlemao
	}

	return prefixo + nome
}
