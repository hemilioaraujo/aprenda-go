package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	t.Run("deve dizer olá para nome específico", func(t *testing.T) {

		resultado := Ola("Manoel")
		esperado := "Olá, Manoel"

		verificaMensagemCorreta(t, resultado, esperado)
	})
	
	t.Run("deve dizer olá mundo quando recebe string vazia", func(t *testing.T) {
		
		resultado := Ola("")
		esperado := "Olá, Mundo"
		
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
