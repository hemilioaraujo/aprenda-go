package arrayeslices

func Soma(numeros []int) (soma int) {
	for _, v := range numeros {
		soma += v
	}

	return
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}

	return
}
