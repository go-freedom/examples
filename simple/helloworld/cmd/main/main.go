package main

import (
	"github.com/8treenet/freedom"
)

func main() {
	// 123
	app := freedom.NewApplication()
	runner := app.NewRunner(":31234")

	app.Iris().Get("/", func(ctx freedom.Context) {
		ctx.WriteString("Hello, world!")
	})

	app.Run(runner, freedom.DefaultConfiguration())
}
