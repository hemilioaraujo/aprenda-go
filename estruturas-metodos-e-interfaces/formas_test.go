package estruturasmetodoseinterfaces

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	ret := Retangulo{10.0, 10.0}
	resultado := Perimetro(ret)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	ret := Retangulo{12.0, 6.0}
	resultado := Area(ret)
	esperado := 72.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}
