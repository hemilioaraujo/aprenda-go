package maps

import (
	"errors"
)

type Dicionario map[string]string

var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra que você procura")

func (d Dicionario) Busca(key string) (string, error) {
	definicao, existe := d[key]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}
