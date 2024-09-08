package iteracao

func Repetir(caractere string, repeticoes int) (resultado string) {
	for i := 0; i < repeticoes; i++ {
		resultado += caractere
	}

	return
}
