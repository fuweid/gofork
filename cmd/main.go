package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = append(app.Commands, cloneCmd)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "gofork: %s\n", err)
		os.Exit(1)
	}
}
