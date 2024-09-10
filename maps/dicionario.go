package maps

type Dicionario map[string]string
type ErrDicionario string

const (
	ErrNaoEncontrado    = ErrDicionario("não foi possível encontrar a palavra que você procura")
	ErrPalavraExistente = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
)

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(key string) (string, error) {
	definicao, existe := d[key]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

func (d Dicionario) Adiciona(key, value string) error {
	_, err := d.Busca(key)
	switch err {
	case ErrNaoEncontrado:
		d[key] = value
	case nil:
		return ErrPalavraExistente
	default:
		return err

	}

	return nil
}
