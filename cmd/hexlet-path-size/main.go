package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"

	"code"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(_ context.Context, command *cli.Command) error {
			path := command.Args().First()
			if path == "" {
				return cli.Exit("path is required", 1)
			}

			isHuman := command.Bool("human")
			isAll := command.Bool("all")
			isRecursive := command.Bool("recursive")

			formatted, err := code.GetPathSize(path, isRecursive, isHuman, isAll)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			fmt.Printf("%s\t%s\n", formatted, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
