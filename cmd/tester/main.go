package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/steelphase/vscode-go-2828/v2/pkg/printer"
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
