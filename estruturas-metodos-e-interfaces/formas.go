package estruturasmetodoseinterfaces

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (r Circulo) Area() float64 {
	return math.Pi * r.Raio * r.Raio
}

func Perimetro(ret Retangulo) float64 {
	return (ret.Largura + ret.Altura) * 2
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}
