package commands

import (
	"fmt"

	"github.com/4strodev/raven/pkg/features/template/domain"
	"github.com/4strodev/raven/pkg/shared/injector"
	"github.com/urfave/cli/v2"
)

var (
	template domain.Template = domain.Template{
		Name: "basic",
		Path: "/home/astro/.local/share/raven/templates/basic",
		Url: "",
	}
)

func GetCreateCommand(inject injector.Injector) *cli.Command {
	return &cli.Command{
		Name: "create",
		Action: func(ctx *cli.Context) error {
			args := ctx.Args()
			if args.Len() != 1 {
				return fmt.Errorf("expected 1 argument provided: %d", args.Len())
			}
			name := args.Get(0)
			_, err := inject.FileSystem.CopyDir(ctx.Context, template.Path, name).Await(ctx.Context)
			if err != nil {
				return err
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "template",
				Aliases: []string{"t"},
				Value:   "",
				Usage:   "Template used to create project",
			},
		},
	}
}

