package mocks

import (
	"fmt"
	"io"
	"time"
)

const (
	inicioContagem = 3
	ultimaPalavra  = "Vai!"
)

// [TODO:] enviar pull request de correção para o repositório pt_BR
// sugerindo a correção dos códigos de exemplo, pois iniciam com o nome Sleep
// e depois passa para Pausa e causa um erro onde deveria funcionar

func Contagem(out io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, ultimaPalavra)
}

type Sleeper interface {
	Sleep()
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

// Pausa vai pausar a execução pela Duração definida
func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}