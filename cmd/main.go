package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "nvim-sync",
		Usage: "Synchronizes neovim configuration between repository and ~/.config/nvim",
		Commands: []*cli.Command{
			{
				Name:  "update",
				Usage: "Sync nvim config from repo to ~/.config/nvim",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return UpdateConfig()
				},
			},
			{
				Name:  "backup",
				Usage: "Sync (backs up) currently installed nvim config from ~/.config/nvim to repo",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return BackupConfig()
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s %v\n", BoldRed("Error:"), err)
		os.Exit(1)
	}
}
