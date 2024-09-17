package main

import (
	"aprenda-go/mocks"
	"os"
	"time"
)

func main() {
	sleeper := &SleeperPadrao{}
	mocks.Contagem(os.Stdout, sleeper)
}

type SleeperPadrao struct{}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}
