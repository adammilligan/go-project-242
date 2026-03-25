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
		Action: func(_ context.Context, command *cli.Command) error {
			path := command.Args().First()
			if path == "" {
				return cli.Exit("path is required", 1)
			}

			size, err := code.GetSize(path)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			fmt.Printf("%dB\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
