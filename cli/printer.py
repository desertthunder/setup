"""ANSI color formatting utilities for terminal output.

Provides simple functions to wrap text with ANSI color codes for use with click.echo.
All functions return colorized strings with reset codes at the end.
"""

import enum


class Colors(enum.StrEnum):
    """ANSI color codes for terminal output."""

    RESET = "\033[0m"
    BOLD = "\033[1m"
    DIM = "\033[2m"

    BLACK = "\033[30m"
    RED = "\033[31m"
    GREEN = "\033[32m"
    YELLOW = "\033[33m"
    BLUE = "\033[34m"
    MAGENTA = "\033[35m"
    CYAN = "\033[36m"
    WHITE = "\033[37m"

    BOLD_BLACK = "\033[1;30m"
    BOLD_RED = "\033[1;31m"
    BOLD_GREEN = "\033[1;32m"
    BOLD_YELLOW = "\033[1;33m"
    BOLD_BLUE = "\033[1;34m"
    BOLD_MAGENTA = "\033[1;35m"
    BOLD_CYAN = "\033[1;36m"
    BOLD_WHITE = "\033[1;37m"


def _colorize(text: str, *colors: Colors) -> str:
    """Apply ANSI color codes to text.

    Args:
        text: Text to colorize
        *colors: Variable number of color codes to apply

    Returns:
        Colorized text with reset code at the end
    """
    color_codes = "".join(colors)
    return f"{color_codes}{text}{Colors.RESET}"


def bold(text: str) -> str:
    """Apply bold formatting to text."""
    return _colorize(text, Colors.BOLD)


def dim(text: str) -> str:
    """Apply dim formatting to text."""
    return _colorize(text, Colors.DIM)


def bold_cyan(text: str) -> str:
    """Apply bold cyan color to text."""
    return _colorize(text, Colors.BOLD_CYAN)


def bold_yellow(text: str) -> str:
    """Apply bold yellow color to text."""
    return _colorize(text, Colors.BOLD_YELLOW)


def bold_magenta(text: str) -> str:
    """Apply bold magenta color to text."""
    return _colorize(text, Colors.BOLD_MAGENTA)


def bold_green(text: str) -> str:
    """Apply bold green color to text."""
    return _colorize(text, Colors.BOLD_GREEN)


def bold_red(text: str) -> str:
    """Apply bold red color to text."""
    return _colorize(text, Colors.BOLD_RED)
