package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/desertthunder/installer/cmd"
	"github.com/urfave/cli/v3"
)

//go:embed packages
var PackageLists embed.FS

func main() {
	root := &cli.Command{
		Name:  "thunderize",
		Usage: "Setup and manage an Arch Linux development environment",
		Commands: []*cli.Command{
			{
				Name:  "install",
				Usage: "Install packages",
				Commands: []*cli.Command{
					{
						Name:  "pacman",
						Usage: "Install packages from official repositories",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.InstallPacmanPackages(PackageLists)
						},
					},
					{
						Name:  "aur",
						Usage: "Install packages from AUR",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.InstallAURPackages(PackageLists)
						},
					},
					{
						Name:  "dev",
						Usage: "Install development tools via asdf",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.InstallDevTools()
						},
					},
					{
						Name:  "all",
						Usage: "Install all packages (pacman, AUR, and dev tools)",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.InstallAllPackages(PackageLists)
						},
					},
				},
			},
			{
				Name:  "config",
				Usage: "Manage configuration files",
				Commands: []*cli.Command{
					{
						Name:  "deploy",
						Usage: "Deploy config(s) from repo to system",
						Arguments: []cli.Argument{
							&cli.StringArg{
								Name:      "name",
								UsageText: "Config name (or 'all' for all configs)",
							},
						},
						Action: func(ctx context.Context, c *cli.Command) error {
							name := c.String("name")
							if name == "" {
								name = "all"
							}
							if name == "all" {
								return cmd.DeployAllConfigs()
							}
							return cmd.DeployConfig(name)
						},
					},
					{
						Name:  "backup",
						Usage: "Backup config(s) from system to repo",
						Arguments: []cli.Argument{
							&cli.StringArg{
								Name:      "name",
								UsageText: "Config name (or 'all' for all configs)",
							},
						},
						Action: func(ctx context.Context, c *cli.Command) error {
							name := c.String("name")
							if name == "" {
								name = "all"
							}
							if name == "all" {
								return cmd.BackupAllConfigs()
							}
							return cmd.BackupConfigByName(name)
						},
					},
					{
						Name:  "list",
						Usage: "List available configurations",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.ListConfigs()
						},
					},
					{
						Name:  "validate",
						Usage: "Validate that all configs exist in repo",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.ValidateConfigs()
						},
					},
				},
			},
			{
				Name:  "setup",
				Usage: "Run full system setup (checks, packages, and configs)",
				Action: func(ctx context.Context, c *cli.Command) error {
					if err := cmd.RunSystemChecks(); err != nil {
						return err
					}

					cmd.Print.Info()
					if err := cmd.InstallAllPackages(PackageLists); err != nil {
						return err
					}

					cmd.Print.Info()
					if err := cmd.DeployAllConfigs(); err != nil {
						return err
					}

					cmd.Print.Info()
					cmd.Print.Success("System setup completed successfully!")
					return nil
				},
			},
			{
				Name:  "check",
				Usage: "Run system checks",
				Action: func(ctx context.Context, c *cli.Command) error {
					return cmd.RunSystemChecks()
				},
			},
			{
				Name:  "secrets",
				Usage: "Manage secrets configuration",
				Commands: []*cli.Command{
					{
						Name:  "init",
						Usage: "Initialize secrets file from template",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.InitSecrets()
						},
					},
					{
						Name:  "check",
						Usage: "Check secrets file existence and permissions",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.CheckSecrets()
						},
					},
					{
						Name:  "edit",
						Usage: "Edit secrets file in default editor",
						Action: func(ctx context.Context, c *cli.Command) error {
							return cmd.EditSecrets()
						},
					},
				},
			},
		},
	}

	if err := root.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s %v\n", cmd.BoldRed("Error:"), err)
		os.Exit(1)
	}
}
