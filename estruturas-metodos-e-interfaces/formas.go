package estruturasmetodoseinterfaces

type Retangulo struct {
	Largura float64
	Altura  float64
}

func Perimetro(ret Retangulo) float64 {
	return (ret.Largura + ret.Altura) * 2
}

func Area(ret Retangulo) float64 {
	return ret.Largura * ret.Altura
}
