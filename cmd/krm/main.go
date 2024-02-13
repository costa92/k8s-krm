package main

import (
	"github.com/costa92/krm/cmd/krm/app"
)

func main() {
	cmd := app.NewAPIServerCommand()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
