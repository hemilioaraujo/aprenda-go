package arrayeslices

// soma todos valores de um slice
func Soma(numeros []int) (soma int) {
	for _, v := range numeros {
		soma += v
	}

	return
}

// soma todos slices
func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}

	return
}

// soma todos os itens de todos slices, exceto o primeiro
// item de cada slice
func SomaTodoOResto(numerosParaSomar ...[]int) (somas []int) {
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			somas = append(somas, Soma(numeros[1:]))
		}
	}

	return
}
