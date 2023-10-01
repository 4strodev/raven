package main

import (
	"fmt"
	"os"

	"github.com/4strodev/raven/pkg/features/project/commands"
	"github.com/4strodev/raven/pkg/shared/injector"
	"github.com/urfave/cli/v2"
)


func main() {
	inject := injector.NewInjector()
	app := &cli.App{
		Name: "raven",
		ExitErrHandler: func(ctx *cli.Context, err error) {
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		Description: "A project creator based on templates",
		Commands: []*cli.Command{
			commands.GetCreateCommand(inject),
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println("hello world")
			return nil
		},
	}

	app.Run(os.Args)
}
