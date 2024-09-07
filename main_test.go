package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	t.Run("deve dizer olá para nome específico", func(t *testing.T) {

		resultado := Ola("Manoel", "")
		esperado := "Olá, Manoel"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("deve dizer olá mundo quando recebe string vazia", func(t *testing.T) {

		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("deve retornar olá em espanhol", func(t *testing.T) {
		resultado := Ola("Alejandro", "espanhol")
		esperado := "Hola, Alejandro"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("deve retornar olá em frances", func(t *testing.T) {
		resultado := Ola("Paul", "frances")
		esperado := "Bonjour, Paul"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
