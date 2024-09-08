package arrayeslices

func Soma(numeros []int) (soma int) {
	for _, v := range numeros {
		soma += v
	}

	return
}
