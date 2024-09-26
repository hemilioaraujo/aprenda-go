package mocks

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	escrita = "escrita"
	pausa   = "pausa"
)

func TestContagem(t *testing.T) {
	t.Run("teste impressão correta", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spyImpressoraSleep := &SpyContagemOperacoes{}

		Contagem(buffer, spyImpressoraSleep)

		resultado := buffer.String()
		esperado := `3
2
1
Vai!`

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("tempo entre impressões", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado %v chamadas, resultado %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})
}

func TestSleeperConfiguravel(t *testing.T) {
    tempoPausa := 5 * time.Second

    tempoSpy := &TempoSpy{}
    sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Pausa}
    sleeper.Pausa()

    if tempoSpy.duracaoPausa != tempoPausa {
        t.Errorf("deveria ter pausado por %v, mas pausou por %v", tempoPausa, tempoSpy.duracaoPausa)
    }
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}
