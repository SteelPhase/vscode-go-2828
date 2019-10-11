package main

import (
	"github.com/caarlos0/env"
	"github.com/steelphase/vscode-go-2828/pkg/printer"
)

type config struct {
	Thing string `env:"THING"`
}

func main() {
	c := config{}

	if err := env.Parse(&c); err != nil {
		panic(err)
	}

	printer.Print(c.Thing)
}
