package ponteiroseerros

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Stringer interface {
	String() string
}

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(valor Bitcoin) error {
	if valor > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= valor

	return nil
}
