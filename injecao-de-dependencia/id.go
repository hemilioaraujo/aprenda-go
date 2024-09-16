package injecaodedependencia

import (
	"fmt"
	"io"
)

func Cumprimenta(w io.Writer, nome string) {
	fmt.Fprintf(w, "Olá, %s", nome)
}
