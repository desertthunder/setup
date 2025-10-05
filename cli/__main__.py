#!/usr/bin/env python3
"""Neovim config sync script.

Updates entire nvim config from repo to neovim or resets from config to repo.
"""

import subprocess
import sys
from pathlib import Path

import click

from . import printer as p


class CommandGroup(click.Group):
    """Custom Click Group that uses ANSI colors for colorful help output."""

    def format_help(self, ctx: click.Context, _: click.HelpFormatter) -> None:
        """Format help output with ANSI colors."""
        click.echo(f"\n{p.bold_cyan('Neovim Config Sync Tool')}\n")
        click.echo(
            f"{p.dim('Synchronizes neovim configuration between this repository and ~/.config/nvim')}\n"
        )
        click.echo(f"{p.bold_yellow('Usage:')}")
        click.echo(f"  python -m cli {p.bold_magenta('[COMMAND]')}\n")
        click.echo(f"{p.bold_yellow('Commands:')}")

        commands = self.list_commands(ctx)
        cmd_list = []
        for cmd_name in commands:
            cmd = self.get_command(ctx, cmd_name)
            if cmd and cmd.help:
                cmd_list.append(f"  {p.bold_magenta(f'{cmd_name:8}')} {cmd.help}")
        click.echo("\n".join(cmd_list) + "\n")


class Command(click.Command):
    """Custom Click Command that uses ANSI colors for help output."""

    def format_help(self, ctx: click.Context, formatter: click.HelpFormatter) -> None:
        """Format command help output with ANSI colors."""
        click.echo(f"\n{p.bold_cyan(self.name)}\n")
        if self.help:
            click.echo(f"{p.dim(self.help)}\n")
        click.echo(f"{p.bold_yellow('Usage:')}")
        click.echo(f"  python -m cli {p.bold_magenta(self.name)}\n")


def get_repo_dir() -> Path:
    """Get the repository directory containing nvim config."""
    script_dir = Path(__file__).parent.parent.resolve()
    return script_dir / "config" / "nvim"


def get_nvim_config() -> Path:
    """Get the nvim config directory path."""
    return Path.home() / ".config" / "nvim"


def run_rsync(source: Path, target: Path, operation: str) -> None:
    """Run rsync to sync config directories.

    Ensures the target directory exists & then builds
    a call to rsync command with sensible args.

    Args:
        source: Source directory to sync from
        target: Target directory to sync to
        operation: Name of the operation for display
    """
    target.mkdir(parents=True, exist_ok=True)

    click.echo(p.bold_cyan(f"{operation} neovim config..."))
    click.echo(f"{p.dim('Source:')} {source}")
    click.echo(f"{p.dim('Target:')} {target}")

    cmd = [
        "rsync",
        "-av",
        "--delete",
        "--exclude=.git",
        "--exclude=*.swp",
        "--exclude=*.swo",
        f"{source}/",
        f"{target}/",
    ]

    try:
        subprocess.run(cmd, check=True, capture_output=True, text=True)
        click.echo(p.bold_green(f"Config {operation.lower()} successfully"))
    except subprocess.CalledProcessError as e:
        click.echo(p.bold_red(f"Error during {operation.lower()}:"), err=True)
        click.echo(e.stderr, err=True)
        sys.exit(1)


@click.group(cls=CommandGroup, invoke_without_command=True)
@click.pass_context
def cli(ctx):
    """Neovim config sync tool.

    Synchronizes neovim configuration between repository and ~/.config/nvim.
    """
    if ctx.invoked_subcommand is None:
        click.echo(ctx.get_help())


@cli.command(cls=Command)
def update():
    """Sync nvim config from repo to ~/.config/nvim"""
    repo_dir = get_repo_dir()
    nvim_config = get_nvim_config()

    if not repo_dir.exists():
        click.echo(
            f"{p.bold_red('Error:')} Repository config not found at {repo_dir}",
            err=True,
        )
        sys.exit(1)

    run_rsync(repo_dir, nvim_config, "Updating")


@cli.command(cls=Command)
def backup():
    """Sync (backs up) currently installed nvim config from ~/.config/nvim to repo"""
    repo_dir = get_repo_dir()
    nvim_config = get_nvim_config()

    if not nvim_config.exists():
        click.echo(
            f"{p.bold_red('Error:')} Nvim config not found at {nvim_config}",
            err=True,
        )
        sys.exit(1)

    run_rsync(nvim_config, repo_dir, "Backing up")


if __name__ == "__main__":
    cli()
